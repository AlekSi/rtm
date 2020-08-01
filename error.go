package rtm

import (
	"fmt"
)

type Error struct {
	Code int    `xml:"code,attr"`
	Msg  string `xml:"msg,attr"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
