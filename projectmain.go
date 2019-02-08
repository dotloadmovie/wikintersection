package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dotloadmovie/wikirelate/network"
)

func main() {
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	arg1Ptr := flag.String("article1", "null", "a string")
	arg2Ptr := flag.String("article2", "null", "a string")

	flag.Parse()

	fmt.Println(*arg1Ptr)
	fmt.Println(*arg2Ptr)

	response := network.GetWiki()

	for _, entry := range response {
		fmt.Printf("%s", entry.Map()["title"])
		fmt.Println("++======++")
	}
}
