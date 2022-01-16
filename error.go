package rtm

import (
	"fmt"
)

// Error represents API error.
type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
