package rtm

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// rtmDuration wraps time.Duration with ISO 8601 (like `PT1H30M`) JSON encoding and decoding,
type rtmDuration struct {
	time.Duration
}

// String implements fmt.Stringer.
func (d rtmDuration) String() string {
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
	}
	return res
}

// MarshalJSON implements json.Marshaler.
func (d rtmDuration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

var durationRE = regexp.MustCompile(`^"PT(\d+H)?(\d+M)?"$`)

// UnmarshalJSON implements json.Unmarshaler.
func (d *rtmDuration) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("rtm.rtmDuration.UnmarshalJSON: too short")
	}

	if len(data) == 2 && data[0] == '"' && data[1] == '"' {
		d.Duration = 0
		return nil
	}

	s := string(data)
	failed := fmt.Errorf("rtm.rtmDuration.UnmarshalJSON: failed to parse %q", s)
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
	_ fmt.Stringer     = rtmDuration{}
	_ json.Marshaler   = rtmDuration{}
	_ json.Unmarshaler = (*rtmDuration)(nil)
)
