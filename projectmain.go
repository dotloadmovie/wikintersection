package main

import (
	"fmt"
	"os"

	"github.com/dotloadmovie/wikintersect/network"
	"github.com/dotloadmovie/wikintersect/utils"
	"github.com/dotloadmovie/wikintersect/view"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	first := network.GetWiki(argsWithoutProg[0])
	second := network.GetWiki(argsWithoutProg[1])

	intersection := utils.GetIntersect(first, second)

	//for _, item := range intersection {
	//	fmt.Printf("%s", item)
	//	fmt.Println("++")
	//}

	view.RenderTable(intersection)
}
