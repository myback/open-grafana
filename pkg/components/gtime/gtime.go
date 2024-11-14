package gtime

import (
	"fmt"
	"strconv"
	"time"
)

const (
	day   = 24 * time.Hour
	week  = 7 * day
	month = 30 * day
	year  = 365 * day
)

// ParseInterval parses an interval with support for all units that Grafana uses.
// An interval is relative to the current wall time.
func ParseInterval(inp string) (time.Duration, error) {
	dur, err := ParseDuration(inp)
	if err != nil {
		return 0, err
	}

	now := time.Now().UTC()
	return now.Add(dur).Sub(now), nil
}

// ParseDuration parses a duration with support for all units that Grafana uses.
// Durations are independent of wall time.
func ParseDuration(inp string) (time.Duration, error) {
	if len(inp) == 0 {
		return 0, nil
	}
	dur, units := inp[:len(inp)-1], inp[len(inp)-1:]
	var mul time.Duration
	switch units {
	case "d":
		mul = day
	case "w":
		mul = week
	case "M":
		mul = month
	case "y":
		mul = year
	default:
		return time.ParseDuration(inp)
	}
	num, err := strconv.Atoi(dur)
	if err != nil {
		return 0, fmt.Errorf("invalid duration %q: %q", inp, err)
	}

	return time.Duration(num) * mul, nil
}
