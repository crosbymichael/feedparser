package feedparser

import (
	"encoding/xml"
	"io"
	"time"
)

const (
	RSS_TIME_FORMAT = time.RFC1123
)

type RssFeed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	PubDate     string `xml:"pubDate"`
	Ttl         int    `xml:"ttl"`
	Item        []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Comments    string `xml:"comments"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Description Body   `xml:"description"`
}

func (i *Item) PubDateTime() (time.Time, error) {
	return time.Parse(RSS_TIME_FORMAT, i.PubDate)
}

func (c *Channel) PubDateTime() (time.Time, error) {
	return time.Parse(RSS_TIME_FORMAT, c.PubDate)
}

func ParseRssFeed(r io.Reader) (*RssFeed, error) {
	dec := xml.NewDecoder(r)
	var f RssFeed
	if err := dec.Decode(&f); err != nil {
		return nil, err
	}
	return &f, nil
}

func WriteRssFeed(f *RssFeed, w io.Writer) error {
	enc := xml.NewEncoder(w)
	return enc.Encode(f)

}
