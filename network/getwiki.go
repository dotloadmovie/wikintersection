package network

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

// GetWiki fetcher for article
func GetWiki(title string) []string {
	links := make([]string, 0)

	combined := requestWiki(title, "initial", links)

	return combined
}

// call against the Wikipedia API and assemble results
func requestWiki(name string, continued string, response []string) []string {
	requestTemplate := []string{
		"https://en.wikipedia.org/w/api.php",
		"?format=json",
		"&formatversion=2",
		"&action=query",
		"&prop=links",
		"&pllimit=200",
	}

	if continued == "null" {
		return response
	}

	requestTemplate = append(requestTemplate, "&titles="+name)

	if continued != "initial" {
		requestTemplate = append(requestTemplate, "&plcontinue="+continued)
	}

	resp, err := http.Get(strings.Join(requestTemplate, ""))

	if err != nil {
		log.Fatal(err)
	}

	jsonResponse, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	raw := parseWiki(string(jsonResponse))

	// not sure why we need this. Cannot just append another array, even though it is of the correct type
	// learnnote: check out what is preventing it working (some kind of type inconsistency?)
	// for the time being, we can use this to cast the slice type to string
	for _, entry := range raw.Array() {
		response = append(response, string(entry.Get("title").String()))
	}

	nextContinued := "null"

	if gjson.Get(string(jsonResponse), "continue.plcontinue").Exists() {
		nextContinued = gjson.Get(string(jsonResponse), "continue.plcontinue").String()
	}

	// recurse through the continue results from Wikipedia
	return requestWiki(name, nextContinued, response)
}

// extract just the links array
func parseWiki(wiki string) gjson.Result {
	return gjson.Get(wiki, "query.pages.0.links")
}
