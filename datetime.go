package rtm

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	dateFormat     = "2006-01-02"
	dateTimeFormat = "2006-01-02T15:04:05Z07:00"
)

// DateTime wraps time.Time with ISO 8601 (like `2006-01-02T15:04:05Z`) JSON encoding and decoding.
type DateTime struct {
	time.Time
	HasTime bool
}

// String implements fmt.Stringer.
func (t DateTime) String() string {
	if t.IsZero() {
		return ""
	}
	// if !t.HasTime {
	// 	return t.Format(dateFormat)
	// }
	return t.Format(dateTimeFormat)
}

// MarshalJSON implements json.Marshaler.
func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *DateTime) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("rtm.DateTime.UnmarshalJSON: too short")
	}

	var tt time.Time
	if len(data) == 2 && data[0] == '"' && data[1] == '"' {
		t.Time = tt
		t.HasTime = false
		return nil
	}

	var err error
	tt, err = time.Parse(`"`+dateTimeFormat+`"`, string(data))
	if err == nil {
		t.Time = tt //.UTC()
		t.HasTime = true
		return nil
	}

	tt, err = time.Parse(`"`+dateFormat+`"`, string(data))
	if err == nil {
		t.Time = tt //.UTC()
		t.HasTime = false
		return nil
	}

	return fmt.Errorf("rtm.DateTime.UnmarshalJSON: %w", err)
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
	_ fmt.Stringer     = DateTime{}
	_ json.Marshaler   = DateTime{}
	_ json.Unmarshaler = (*DateTime)(nil)
)
