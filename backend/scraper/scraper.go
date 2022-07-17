package main

import (
	"github.com/NERSAION/hari/backend/secrets"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
)

func main() {
	var text string
	secrets, err := secrets.ReadSecrets("scraper")
	if err != nil {
		log.Fatalln("cant parse secrets: ", err)
	}
	c := colly.NewCollector(
		colly.AllowedDomains(secrets["domain"]),
	)
	c.OnHTML(".entry-content", func(e *colly.HTMLElement) {
		text = e.Text
	})
	c.Visit(secrets["url"])
	quotes := strings.Split(text, secrets["sep"])
	for n, q := range quotes {
		for _, l := range strings.Split(q, "\n") {
			if len(l) > 10 {
				quotes[n] = l
				break
			}
		}
	}
	// todo: write to database
}
