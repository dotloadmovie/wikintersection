package network

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// GetSearch sends a string to the Wiki API and responds with a list of matching articles
func GetSearch(search string) []string {
	links := make([]string, 0)

	combined := requestWikiSearch(search, "0", links)

	return combined
}

func requestWikiSearch(name string, continued string, response []string) []string {
	requestTemplate := []string{
		"https://en.wikipedia.org/w/api.php",
		"?action=query",
		"&srlimit=300",
		"&list=search",
		"&utf8=&format=json",
	}

	if continued == "null" {
		return response
	}

	requestTemplate = append(requestTemplate, "&srsearch=intitle:"+name)

	requestTemplate = append(requestTemplate, "&sroffset="+continued)

	resp, err := http.Get(strings.Join(requestTemplate, ""))

	if err != nil {
		log.Fatal(err)
	}

	jsonResponse, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	raw := parseWikiSearch(string(jsonResponse))

	// not sure why we need this. Cannot just append another array, even though it is of the correct type
	// learnnote: check out what is preventing it working (some kind of type inconsistency?)
	// for the time being, we can use this to cast the slice type to string
	for _, entry := range raw.Array() {
		response = append(response, string(entry.Get("title").String()))
	}

	nextContinued := "null"

	if gjson.Get(string(jsonResponse), "continue.sroffset").Exists() {
		nextContinued = gjson.Get(string(jsonResponse), "continue.sroffset").String()
	}

	// recurse through the continue results from Wikipedia
	return requestWikiSearch(name, nextContinued, response)
}

// extract just the links array
func parseWikiSearch(wiki string) gjson.Result {
	return gjson.Get(wiki, "query.search")
}
