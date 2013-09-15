package feedparser

import (
	"bytes"
	"os"
	"testing"
)

func TestParseRss(t *testing.T) {
	f, err := os.Open("test_rss.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	feed, err := ParseRssFeed(f)
	if err != nil {
		t.Fatal(err)
	}

	if feed == nil {
		t.FailNow()
	}

	if feed.Channel.Title != "Hacker News" {
		t.FailNow()
	}
	if feed.Version != "2.0" {
		t.FailNow()
	}
	if len(feed.Channel.Item) != 30 {
		t.FailNow()
	}
}

func TestParseFeedburner(t *testing.T) {
	f, err := os.Open("test_rss_ars.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	feed, err := ParseRssFeed(f)
	if err != nil {
		t.Fatal(err)
	}

	if feed == nil {
		t.FailNow()
	}

	if feed.Channel.Title != "Ars Technica" {
		t.FailNow()
	}
	if feed.Version != "2.0" {
		t.FailNow()
	}

	if feed.Channel.Ttl != 30 {
		t.FailNow()
	}
	if len(feed.Channel.Item) != 25 {
		t.FailNow()
	}

	if _, err := feed.Channel.PubDateTime(); err != nil {
		t.Fatal(err)
	}
	item := feed.Channel.Item[0]
	if _, err := item.PubDateTime(); err != nil {
		t.Fatal(err)
	}
}

func TestWriteRssFeed(t *testing.T) {
	f, err := os.Open("test_rss.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	feed, err := ParseRssFeed(f)
	if err != nil {
		t.Fatal(err)
	}

	if feed == nil {
		t.FailNow()
	}

	if feed.Channel.Title != "Hacker News" {
		t.FailNow()
	}
	if feed.Version != "2.0" {
		t.FailNow()
	}
	if len(feed.Channel.Item) != 30 {
		t.FailNow()
	}

	var w bytes.Buffer
	if err := WriteRssFeed(feed, &w); err != nil {
		t.Fatal(err)
	}
	if w.Len() < 1024 {
		t.Fail()
	}
}
