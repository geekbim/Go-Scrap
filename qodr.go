package main

import (
	"github.com/gocolly/colly"
    "fmt"
	"log"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Prinsip struct {
	Title			string 	`json:"title"`
	Description		string	`json:"desc"`
}

func main() {
	
	allPrinsips := make([]Prinsip, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.qodr.or.id", "qodr.or.id"),
	)

	collector.OnHTML(".rail-cards .card-how-to .inner-content-cards", func(element *colly.HTMLElement) {
		prinsipTitle := element.ChildText("h5")
		prinsipDesc := element.ChildText("div")

		prinsip := Prinsip {
			Title:			prinsipTitle,
			Description:	prinsipDesc,
		}

		allPrinsips = append(allPrinsips, prinsip)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})

	collector.Visit("http://qodr.or.id/")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allPrinsips)

	writeJSON(allPrinsips)

}

func writeJSON(data []Prinsip) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}

	ioutil.WriteFile("scrap.json", dataFile, 0666)
}