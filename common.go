package feedparser

type Body struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}
