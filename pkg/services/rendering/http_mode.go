package rendering

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/myback/open-grafana/pkg/setting"
)

var netTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	Dial: (&net.Dialer{
		Timeout: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 5 * time.Second,
}

var netClient = &http.Client{
	Transport: netTransport,
}

func (rs *RenderingService) renderViaHttp(ctx context.Context, renderKey string, opts Opts) (*RenderResult, error) {
	filePath, err := rs.getFilePathForNewImage()
	if err != nil {
		return nil, err
	}

	rendererUrl, err := url.Parse(rs.Cfg.RendererUrl)
	if err != nil {
		return nil, err
	}

	queryParams := rendererUrl.Query()
	queryParams.Add("url", rs.getURL(opts.Path))
	queryParams.Add("renderKey", renderKey)
	queryParams.Add("width", strconv.Itoa(opts.Width))
	queryParams.Add("height", strconv.Itoa(opts.Height))
	queryParams.Add("domain", rs.domain)
	queryParams.Add("timezone", isoTimeOffsetToPosixTz(opts.Timezone))
	queryParams.Add("encoding", opts.Encoding)
	queryParams.Add("timeout", strconv.Itoa(int(opts.Timeout.Seconds())))
	queryParams.Add("deviceScaleFactor", fmt.Sprintf("%f", opts.DeviceScaleFactor))
	rendererUrl.RawQuery = queryParams.Encode()

	req, err := http.NewRequest("GET", rendererUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("Grafana/%s", setting.BuildVersion))

	for k, v := range opts.Headers {
		req.Header[k] = v
	}

	// gives service some additional time to timeout and return possible errors.
	reqContext, cancel := context.WithTimeout(ctx, opts.Timeout+time.Second*2)
	defer cancel()

	req = req.WithContext(reqContext)

	rs.log.Debug("calling remote rendering service", "url", rendererUrl)

	// make request to renderer server
	resp, err := netClient.Do(req)
	if err != nil {
		rs.log.Error("Failed to send request to remote rendering service.", "error", err)
		return nil, fmt.Errorf("failed to send request to remote rendering service: %w", err)
	}

	// save response to file
	defer func() {
		if err := resp.Body.Close(); err != nil {
			rs.log.Warn("Failed to close response body", "err", err)
		}
	}()

	// check for timeout first
	if errors.Is(reqContext.Err(), context.DeadlineExceeded) {
		rs.log.Info("Rendering timed out")
		return nil, ErrTimeout
	}

	// if we didn't get a 200 response, something went wrong.
	if resp.StatusCode != http.StatusOK {
		rs.log.Error("Remote rendering request failed", "error", resp.Status)
		return nil, fmt.Errorf("remote rendering request failed, status code: %d, status: %s", resp.StatusCode,
			resp.Status)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := out.Close(); err != nil {
			// We already close the file explicitly in the non-error path, so shouldn't be a problem
			rs.log.Warn("Failed to close file", "path", filePath, "err", err)
		}
	}()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		// check that we didn't timeout while receiving the response.
		if errors.Is(reqContext.Err(), context.DeadlineExceeded) {
			rs.log.Info("Rendering timed out")
			return nil, ErrTimeout
		}
		rs.log.Error("Remote rendering request failed", "error", err)
		return nil, fmt.Errorf("remote rendering request failed: %w", err)
	}
	if err := out.Close(); err != nil {
		return nil, fmt.Errorf("failed to write to %q: %w", filePath, err)
	}

	return &RenderResult{FilePath: filePath}, err
}
