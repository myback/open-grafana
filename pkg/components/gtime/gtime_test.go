package gtime

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseInterval(t *testing.T) {
	tcs := []struct {
		inp      string
		duration time.Duration
		err      *regexp.Regexp
	}{
		{inp: "1d", duration: day},
		{inp: "1w", duration: week},
		{inp: "2w", duration: 2 * week},
		{inp: "1M", duration: month},
		{inp: "1y", duration: year},
		{inp: "5y", duration: 5 * year},
		{inp: "invalid-duration", err: regexp.MustCompile(`^time: invalid duration "?invalid-duration"?$`)},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			res, err := ParseInterval(tc.inp)
			if tc.err == nil {
				require.NoError(t, err, "input %q", tc.inp)
				require.Equal(t, tc.duration, res, "input %q", tc.inp)
			} else {
				require.Error(t, err, "input %q", tc.inp)
				require.Regexp(t, tc.err, err.Error())
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	tcs := []struct {
		inp      string
		duration time.Duration
		err      *regexp.Regexp
	}{
		{inp: "1s", duration: time.Second},
		{inp: "1m", duration: time.Minute},
		{inp: "1h", duration: time.Hour},
		{inp: "1d", duration: day},
		{inp: "1w", duration: week},
		{inp: "2w", duration: 2 * week},
		{inp: "1M", duration: month},
		{inp: "1y", duration: year},
		{inp: "5y", duration: 5 * year},
		{inp: "invalid-duration", err: regexp.MustCompile(`^time: invalid duration "?invalid-duration"?$`)},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			res, err := ParseDuration(tc.inp)
			if tc.err == nil {
				require.NoError(t, err, "input %q", tc.inp)
				require.Equal(t, tc.duration, res, "input %q", tc.inp)
			} else {
				require.Error(t, err, "input %q", tc.inp)
				require.Regexp(t, tc.err, err.Error())
			}
		})
	}
}
