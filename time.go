package rtm

import (
	"encoding"
	"fmt"
	"time"
)

// Time wraps time.Time with ISO 8601 (like `2019-01-20T09:20:58Z`) text encoding and decoding.
type Time struct {
	time.Time
}

func (t Time) String() string {
	return t.UTC().Format(time.RFC3339)
}

// MarshalText implements encoding.TextMarshaler.
func (t Time) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (t *Time) UnmarshalText(data []byte) error {
	var tt time.Time
	if len(data) == 0 {
		t.Time = tt
		return nil
	}

	var err error
	if tt, err = time.Parse(time.RFC3339, string(data)); err != nil {
		return err
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
	_ fmt.Stringer             = Time{}
	_ encoding.TextMarshaler   = Time{}
	_ encoding.TextUnmarshaler = (*Time)(nil)
)
