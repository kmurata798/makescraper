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
	Content []string `json:"content"`
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

	var r []string
	// On every a element which has href attribute call callback
	// wikipedia new featured article selector
	c.OnHTML("#mp-tfa > p", func(e *colly.HTMLElement) {

		// Print link
		r = append(r, e.Text)
		data1 := scrapedData{Content: r}
		fmt.Println(data1)

		// scrapedJSON, _ := json.Marshal(data1)
		scrapedJSON, _ := json.MarshalIndent(data1, "", "    ")
		fmt.Println(string(scrapedJSON))

	})
	// wikipedia Recent News
	c.OnHTML("#mp-itn", func(e *colly.HTMLElement) {

		// Print link
		r = append(r, e.Text)
		data2 := scrapedData{Content: r}
		fmt.Println(data2)

		// scrapedJSON2, _ := json.Marshal(data2)
		scrapedJSON, _ := json.MarshalIndent(data2, "", "    ")
		fmt.Println(string(scrapedJSON))

		writeFile(arg, string(scrapedJSON))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://en.wikipedia.org/wiki/Main_Page")

}
