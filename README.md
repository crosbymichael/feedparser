## Go Feedparser

RSS 2.0 and Atom feed parser for Go

## Example
```golang
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
```

### License
MIT

Copyright (c) 2013 Michael Crosby. michael@crosbymichael.com

Permission is hereby granted, free of charge, to any person
obtaining a copy of this software and associated documentation 
files (the "Software"), to deal in the Software without 
restriction, including without limitation the rights to use, copy, 
modify, merge, publish, distribute, sublicense, and/or sell copies 
of the Software, and to permit persons to whom the Software is 
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be 
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. 
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT 
HOLDERS BE LIABLE FOR ANY CLAIM, 
DAMAGES OR OTHER LIABILITY, 
WHETHER IN AN ACTION OF CONTRACT, 
TORT OR OTHERWISE, 
ARISING FROM, OUT OF OR IN CONNECTION WITH 
THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
