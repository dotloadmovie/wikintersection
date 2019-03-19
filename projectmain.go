package main

import (
	"fmt"
	"github.com/dotloadmovie/wikintersect/network"
	"github.com/dotloadmovie/wikintersect/utils"
	"github.com/dotloadmovie/wikintersect/view"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	switchCommand := argsWithoutProg[0]

	if switchCommand == "compare" {
		getIntersection(argsWithoutProg[1], argsWithoutProg[2])
	}

	if switchCommand == "search" {
		getSearch(argsWithoutProg[1])
	}

}

// getIntersection: get the intersection of two articles from Wikipedia
func getIntersection(first string, second string) {
	firstResults := network.GetWiki(first)
	secondResults := network.GetWiki(second)

	intersection := utils.GetIntersect(firstResults, secondResults)

	view.RenderTable(intersection)
}

// getSearch: get a list of article matches from Wikipedia
func getSearch(searchString string) {
	results := network.GetSearch(searchString)
	view.RenderTable(results)
}
