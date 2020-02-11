package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type scrapedData struct {
	content string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#mp-tfa > p", func(e *colly.HTMLElement) {

		// Print link
		data := scrapedData{content: e.Text}
		fmt.Println(data)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://en.wikipedia.org/wiki/Main_Page")

}
