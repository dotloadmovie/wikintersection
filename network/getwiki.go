package network

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

// GetWiki fetcher for article
func GetWiki() []gjson.Result {
	links := make([]gjson.Result, 0)

	combined := requestWiki("Albert%20Einstein", "null", links)

	return combined
}

func requestWiki(name string, continued string, response []gjson.Result) []gjson.Result {
	resp, err := http.Get("https://en.wikipedia.org/w/api.php?format=json&formatversion=2&action=query&titles=Albert%20Einstein&prop=links&pllimit=500")

	if err != nil {
		log.Fatal(err)
	}

	jsonResponse, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	raw := parseWiki(string(jsonResponse))

	for _, entry := range raw.Array() {
		response = append(response, entry)
	}

	return response
}

func parseWiki(wiki string) gjson.Result {
	return gjson.Get(wiki, "query.pages.0.links")
}
