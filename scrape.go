package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gocolly/colly"
)

// Struct that can be used for json
type scrapedData struct {
	Content string `json:"content"`
}

func writeFile(name string, data string) {
	/*
		Makesite MVP

			Writes data onto file
	*/
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	arg := os.Args[1]
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#mp-tfa > p", func(e *colly.HTMLElement) {

		// Print link
		data := scrapedData{Content: e.Text}
		fmt.Println(data)

		// scrapedJSON, _ := json.Marshal(data)
		scrapedJSON1, _ := json.MarshalIndent(data, "", "    ")
		fmt.Println(string(scrapedJSON1))

		writeFile(arg, string(scrapedJSON1))
	})

	c.OnHTML("#mp-itn", func(e *colly.HTMLElement) {

		// Print link
		data := scrapedData{Content: e.Text}
		fmt.Println(data)

		// scrapedJSON, _ := json.Marshal(data)
		scrapedJSON2, _ := json.MarshalIndent(data, "", "    ")
		fmt.Println(string(scrapedJSON2))

		// writeFile(arg, string(scrapedJSON2))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://en.wikipedia.org/wiki/Main_Page")

}
