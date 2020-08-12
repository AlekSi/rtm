package rtm

import (
	"encoding/json"
	"fmt"
	"time"
)

// Time wraps time.Time with ISO 8601 (like `2019-01-20T09:20:58Z`) JSON encoding and decoding.
type Time struct {
	time.Time
}

func (t Time) String() string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format(time.RFC3339)
}

// MarshalJSON implements json.Marshaler.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Time) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("rtm.Time.UnmarshalJSON: too short")
	}

	var tt time.Time
	if len(data) == 2 && data[0] == '"' && data[1] == '"' {
		t.Time = tt
		return nil
	}

	var err error
	if tt, err = time.Parse(`"`+time.RFC3339+`"`, string(data)); err != nil {
		return fmt.Errorf("rtm.Time.UnmarshalJSON: %w", err)
	}

	t.Time = tt.UTC()
	return nil
}

// // hasTime returns false if this instance contains only date without time.
// func (t Time) hasTime() bool {
// 	hour, min, sec := t.Clock()
// 	return hour != 0 || min != 0 || sec != 0
// }

// func (t Time) withoutTime() Time {
// 	year, month, day := t.Date()
// 	return Time{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
// }

// check interfaces
var (
	_ fmt.Stringer     = Time{}
	_ json.Marshaler   = Time{}
	_ json.Unmarshaler = (*Time)(nil)
)
