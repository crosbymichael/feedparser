package feedparser

import (
	"bytes"
	"os"
	"testing"
)

func TestParseAtom(t *testing.T) {
	f, err := os.Open("test_atom.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	feed, err := ParseAtomFeed(f)
	if err != nil {
		t.Fatal(err)
	}

	if feed == nil {
		t.FailNow()
	}

	if feed.Title != "Michael Crosby" {
		t.FailNow()
	}

	if len(feed.Entry) != 33 {
		t.FailNow()
	}
}

func TestWriteAtomFeed(t *testing.T) {
	f, err := os.Open("test_atom.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	feed, err := ParseAtomFeed(f)
	if err != nil {
		t.Fatal(err)
	}

	if feed == nil {
		t.FailNow()
	}

	if feed.Title != "Michael Crosby" {
		t.FailNow()
	}

	if len(feed.Entry) != 33 {
		t.FailNow()
	}

	var w bytes.Buffer
	if err := WriteAtomFeed(feed, &w); err != nil {
		t.Fatal(err)
	}
	if w.Len() < 1024 {
		t.Fail()
	}
}
