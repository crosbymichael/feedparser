package main

import (
	"fmt"
	"github.com/crosbymichael/feedparser"
	"net/http"
	"time"
)

func main() {
	r, err := http.Get("http://crosbymichael.com/feeds/all.atom.xml")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	feed, err := feedparser.ParseAtomFeed(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Feed: %s\nTitle: %s\n", feed.Link[0].Href, feed.Title)

	for _, e := range feed.Entry {
		fmt.Printf("Title: %s\nUpdated: %s\n", e.Title, e.Updated.Format(time.RubyDate))
	}
}
