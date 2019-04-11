package rtm

import (
	"encoding"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Duration wraps time.Duration with ISO 8601 (like `PT1H30M`) text encoding and decoding,
type Duration struct {
	time.Duration
}

func (d Duration) String() string {
	panic("rtm.Duration.String is not implemented!")
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

// check interfaces
var (
	_ fmt.Stringer             = Duration{}
	_ encoding.TextMarshaler   = Duration{}
	_ encoding.TextUnmarshaler = (*Duration)(nil)
)
