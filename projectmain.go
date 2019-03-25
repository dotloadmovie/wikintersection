package main

import (
	"fmt"
	"github.com/dotloadmovie/wikintersection/network"
	"github.com/dotloadmovie/wikintersection/utils"
	"github.com/dotloadmovie/wikintersection/view"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	switchCommand := argsWithoutProg[0]

	if switchCommand == "compare" {
		getIntersection(argsWithoutProg[1], argsWithoutProg[2], map[string]string {
			"format": "json",
			"formatversion": "2",
			"action": "query",
			"prop": "links",
			"pllimit": "200",
		})
	}

	if switchCommand == "search" {
		getSearch(argsWithoutProg[1], map[string]string {
			"action": "query",
			"srlimit": "300",
			"list": "search",
			"&utf8":"",
			"format":"json",
		})
	}

}

// getIntersection: get the intersection of two articles from Wikipedia
func getIntersection(first string, second string, params map[string]string) {
	network.InitWiki(params)

	firstResults := network.GetWiki(first)
	secondResults := network.GetWiki(second)

	intersection := utils.GetIntersect(firstResults, secondResults)

	view.RenderTable(intersection)
}

// getSearch: get a list of article matches from Wikipedia
func getSearch(searchString string, params map[string]string) {
	network.InitSearch(params)

	results := network.GetSearch(searchString)
	view.RenderTable(results)
}
