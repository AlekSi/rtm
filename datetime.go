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
	hasTime bool // do not set this field directly
}

func (t *DateTime) stripTime() {
	t.Time = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t.hasTime = false
}

// String implements fmt.Stringer.
func (t DateTime) String() string {
	if t.IsZero() {
		return ""
	}
	if t.DateOnly() {
		return t.Format(dateFormat)
	}
	return t.Format(dateTimeFormat)
}

// DateOnly returns true if this DateTime has no time part.
func (t DateTime) DateOnly() bool {
	return !t.hasTime
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
		t.hasTime = false
		return nil
	}

	var err error
	tt, err = time.Parse(`"`+dateTimeFormat+`"`, string(data))
	if err == nil {
		t.Time = tt
		t.hasTime = true
		return nil
	}

	tt, err = time.Parse(`"`+dateFormat+`"`, string(data))
	if err == nil {
		t.Time = tt
		t.hasTime = false
		return nil
	}

	return fmt.Errorf("rtm.DateTime.UnmarshalJSON: %w", err)
}

// check interfaces
var (
	_ fmt.Stringer     = DateTime{}
	_ json.Marshaler   = DateTime{}
	_ json.Unmarshaler = (*DateTime)(nil)
)
