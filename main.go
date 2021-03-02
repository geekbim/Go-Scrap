package main

import (
	"github.com/gocolly/colly"
	"strconv"
    "fmt"
	"log"
	"encoding/json"
	"os"
	"io/ioutil"
)

// Info of a struct for website data
type Info struct {
	ID 			int 	`json:"id"`
	Description string 	`json:"description"`
}

func main() {
	
	allInfos := make([]Info, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("www.factretriever.com", "factretriever.com"),
	)

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		infoID, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Error : ", err)
		}

		infoDesc := element.Text

		info := Info {
			ID:				infoID,
			Description:	infoDesc,
		}

		allInfos = append(allInfos, info)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting ", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/evolution-facts")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allInfos)

	writeJSON(allInfos)

}

func writeJSON(data []Info) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}

	ioutil.WriteFile("scrap.json", dataFile, 0666)
}