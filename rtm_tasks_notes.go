package rtm

type Note struct {
	ID       string `xml:"id,attr"`
	Created  Time   `xml:"created,attr"`
	Modified Time   `xml:"modified,attr"`
	Text     string `xml:",chardata"`
}
