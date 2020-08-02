package rtm

import (
	"fmt"
)

type Error struct {
	Code int    `xml:"code,attr" json:"code"`
	Msg  string `xml:"msg,attr" json:"msg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
