package rtm

import (
	"encoding/json"
	"fmt"
)

// rtmBool wraps bool with JSON string encoding and decoding.
type rtmBool bool

// MarshalJSON implements json.Marshaler.
func (b rtmBool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte(`"1"`), nil
	}
	return []byte(`"0"`), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (b *rtmBool) UnmarshalJSON(data []byte) error {
	switch s := string(data); s {
	case `"0"`:
		*b = false
	case `"1"`:
		*b = true
	default:
		return fmt.Errorf("rtm.rtmBool.UnmarshalJSON: unexpected %q", s)
	}
	return nil
}

// check interfaces
var (
	_ json.Marshaler   = rtmBool(true)
	_ json.Unmarshaler = (*rtmBool)(nil)
)
