package rtm

import (
	"encoding/xml"
)

type Transaction struct {
	XMLName  xml.Name `xml:"transaction"`
	ID       string   `xml:"id,attr"`
	Undoable bool     `xml:"undoable,attr"`
}
