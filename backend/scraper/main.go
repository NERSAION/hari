package main

import (
	"fmt"
	"github.com/NERSAION/hari/backend/secrets"
	"github.com/gocolly/colly/v2"
	"log"
)

func main() {
	secrets, err := secrets.ReadSecrets("scraper")
	if err != nil {
		log.Fatalln("cant parse secrets: ", err)
	}
	c := colly.NewCollector(
		colly.AllowedDomains(secrets["domain"]),
	)
	c.OnHTML("entry-content", func(e *colly.HTMLElement) {
		ps := e.ChildAttrs("", "p")
		fmt.Println(ps)
	})
}
