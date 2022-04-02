package rtm

import (
	"fmt"
)

// Error represents API error.
type Error struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

// check interfaces
var (
	_ error = (*Error)(nil)
)
