package rtm

import "encoding/xml"

type transaction struct {
	XMLName  xml.Name `xml:"transaction"`
	ID       string   `xml:",attr"`
	Undoable bool     `xml:",attr"`
}
