package feedparser

import (
	"encoding/xml"
	"io"
	"time"
)

type AtomFeed struct {
	XMLName xml.Name  `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated,attr"`
	Author  Author    `xml:"author"`
	Entry   []Entry   `xml:"entry"`
}

type Entry struct {
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated"`
	Author  Author    `xml:"author"`
	Summary Body      `xml:"summary"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

type Author struct {
	Name  string `xml:"name"`
	URI   string `xml:"uri"`
	Email string `xml:"email"`
}

func ParseAtomFeed(r io.Reader) (*AtomFeed, error) {
	dec := xml.NewDecoder(r)

	var f AtomFeed
	if err := dec.Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}

// Write the atom feed to the underlying writer
func WriteAtomFeed(f *AtomFeed, w io.Writer) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(f)
}
