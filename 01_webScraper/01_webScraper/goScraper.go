package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
	// filename for data
	fName := "data.csv"
	// create a file
	file, err := os.Create(fName)
	// check for errors
	if err != nil {
		log.Fatalf("Could not create file, error : %q", err)
		return
	}
	// close file afterwards
	defer file.Close()

	// instantiate a csv writer
	writer := csv.NewWriter(file)
	// flush contents afterwards
	defer writer.Flush()

	// instantiate a collector
	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	// point to the webpage structure you need to fetch
	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		// write the desired data into csv
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	// loop through all the pages of the website
	for i := 0; i < 286; i++ {
		fmt.Printf("Scraping Page : %d\n", i)
		// visit each page and scrape data
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	// let yourself know when you're done :D
	log.Printf("Scraping Finished\n")
	log.Println(c)
}
