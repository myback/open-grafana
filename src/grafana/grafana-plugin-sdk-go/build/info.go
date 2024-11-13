package build

import (
	"encoding/json"
	"fmt"
	"time"
)

// set from -X
var buildInfoJSON string

// exposed for testing.
var now = time.Now

// Info See also PluginBuildInfo in https://github.com/myback/grafana/blob/master/pkg/plugins/models.go
type Info struct {
	Time    int64  `json:"time,omitempty"`
	Version string `json:"version,omitempty"`
}

// this will append build flags -- the keys are picked to match existing
// grafana build flags from bra
func (v Info) appendFlags(flags map[string]string) {
	if v.Version != "" {
		flags["main.version"] = v.Version
	}

	out, err := json.Marshal(v)
	if err == nil {
		flags["github.com/grafana/grafana-plugin-sdk-go/build.buildInfoJSON"] = string(out)
	}
}

// GetBuildInfo returns the build information that was compiled into the binary using:
// -X `github.com/grafana/grafana-plugin-sdk-go/build.buildInfoJSON={...}`
func GetBuildInfo() (Info, error) {
	v := Info{}
	if buildInfoJSON == "" {
		return v, fmt.Errorf("build info was now set when this was compiled")
	}
	err := json.Unmarshal([]byte(buildInfoJSON), &v)
	return v, err
}
