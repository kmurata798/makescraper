package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

// Struct that can be used for json
type scrapedData struct {
	gorm.Model
	Content []string `json:"content"`
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

// Function that runs the Echo server
func startServer(dataSlice []string) {
	e := echo.New()

	e.GET("/", func(f echo.Context) error {
		return f.JSON(http.StatusOK, dataSlice)
	})

	fmt.Println("Server running: http://localhost:8000")
	e.Logger.Fatal(e.Start(":8000"))
}

func main() {
	// Gorm Attempt --> error in the console:
	// panic: invalid sql type  (slice) for sqlite3

	// goroutine 1 [running]:
	// github.com/jinzhu/gorm.(*sqlite3).DataTypeOf(0xc00021d620, 0xc000175680, 0xc, 0x4805518)
	// 		/Users/kento/go/pkg/mod/github.com/jinzhu/gorm@v1.9.12/dialect_sqlite3.go:64 +0x751
	// github.com/jinzhu/gorm.(*Scope).createTable(0xc000222480, 0xc000200bf0)
	// 		/Users/kento/go/pkg/mod/github.com/jinzhu/gorm@v1.9.12/scope.go:1169 +0x27c
	// github.com/jinzhu/gorm.(*Scope).autoMigrate(0xc000222480, 0x46fa300)
	// 		/Users/kento/go/pkg/mod/github.com/jinzhu/gorm@v1.9.12/scope.go:1265 +0x400
	// github.com/jinzhu/gorm.(*DB).AutoMigrate(0xc000217790, 0xc000135f40, 0x1, 0x1, 0x1)
	// 		/Users/kento/go/pkg/mod/github.com/jinzhu/gorm@v1.9.12/main.go:684 +0x96
	// main.main()
	// 		/Users/kento/go/src/makescraper/scrape.go:69 +0x15c

	// db, err := gorm.Open("sqlite3", "wiki-article.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// Slice declared to hold the strings I have scraped
	var dataSlice []string

	// // Migrate the schema
	// db.AutoMigrate(&scrapedData{})

	// // Create
	// db.Create(&scrapedData{Content: dataSlice})

	// var dbData scrapedData
	// db.First(&dbData, 1) // find product with id 1
	// db.First(&dbData, ) // find product with code l1212

	// // Update - update product's price to 2000
	// db.Model(&dbData).Update("Price", 2000)

	// // Delete - delete product
	// db.Delete(&dbData)

	// User input for file I want to write my scraped data (strings) into --> output.json
	arg := os.Args[1]
	// Instantiate default collector
	c := colly.NewCollector()

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

	// Starts the server
	startServer(dataSlice)

}

// How to run program:
// $ go build
// $ ./makescraper output.json
