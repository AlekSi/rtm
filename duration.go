package rtm

import (
	"encoding"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Duration wraps time.Duration with ISO 8601 (like `PT1H30M`) text encoding and decoding,
type Duration struct {
	time.Duration
}

func (d Duration) String() string {
	if d.Duration == 0 {
		return ""
	}

	res := "PT"
	td := d.Duration
	if h := td / time.Hour; h > 0 {
		res += strconv.Itoa(int(h)) + "H"
		td -= h * time.Hour
	}
	if m := td / time.Minute; m > 0 {
		res += strconv.Itoa(int(m)) + "M"
		td -= m * time.Minute
	}
	if s := td / time.Second; s > 0 {
		res += strconv.Itoa(int(s)) + "S"
	}
	return res
}

// MarshalText implements encoding.TextMarshaler.
func (d Duration) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

var durationRE = regexp.MustCompile(`^PT(\d+H)?(\d+M)?$`)

// UnmarshalText implements encoding.TextUnmarshaler.
func (d *Duration) UnmarshalText(data []byte) error {
	if len(data) == 0 {
		d.Duration = time.Duration(0)
		return nil
	}

	s := string(data)
	failed := fmt.Errorf("rtm.Duration.UnmarshalText: failed to parse %q", s)
	matches := durationRE.FindStringSubmatch(s)
	if len(matches) != 3 {
		return failed
	}

	hs, ms := strings.ToLower(matches[1]), strings.ToLower(matches[2])
	var h, m time.Duration
	var err error
	if len(hs) != 0 {
		if h, err = time.ParseDuration(hs); err != nil {
			return failed
		}
	}
	if len(ms) != 0 {
		if m, err = time.ParseDuration(ms); err != nil {
			return failed
		}
	}

	d.Duration = h + m
	return nil
}

// parseDuration parses duration from string ignoring errors.
// See Duration.UnmarshalText for the version with error handling.
func parseDuration(s string) Duration {
	var d Duration
	_ = d.UnmarshalText([]byte(s))
	return d
}

// check interfaces
var (
	_ fmt.Stringer             = Duration{}
	_ encoding.TextMarshaler   = Duration{}
	_ encoding.TextUnmarshaler = (*Duration)(nil)
)
