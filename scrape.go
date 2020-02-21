package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/labstack/echo"
)

// Struct that can be used for json
type scrapedData struct {
	Content []string `json:"content"`
}

type jsonData struct {
	message []string `json:"message"`
}

func writeFile(name string, data string) {
	/*
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

func readFile(name string) string {
	/*
		Collects data from file
	*/
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(fileContents)

}

func startServer(dataSlice []string) {
	e := echo.New()

	e.GET("/", func(f echo.Context) error {
		return f.JSON(http.StatusOK, dataSlice)
	})

	fmt.Println("Server running: http://localhost:8000")
	e.Logger.Fatal(e.Start(":8000"))
}

func main() {
	arg := os.Args[1]
	// Instantiate default collector
	c := colly.NewCollector()

	var dataSlice []string
	// On every a element which has href attribute call callback
	// wikipedia new featured article selector
	c.OnHTML("#mp-tfa > p", func(e *colly.HTMLElement) {

		// Print link
		dataSlice = append(dataSlice, e.Text)
		data1 := scrapedData{Content: dataSlice}
		fmt.Println(data1)

		// scrapedJSON, _ := json.Marshal(data1)
		scrapedJSON, _ := json.MarshalIndent(data1, "", "    ")
		fmt.Println(string(scrapedJSON))

	})
	// wikipedia Recent News
	c.OnHTML("#mp-itn", func(e *colly.HTMLElement) {

		// Print link
		dataSlice = append(dataSlice, e.Text)
		data2 := scrapedData{Content: dataSlice}
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

	// fileContents := readFile(arg)
	startServer(dataSlice)

}
