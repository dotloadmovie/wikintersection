package view

import (
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// RenderTable shows results as a numerical table view
func RenderTable(resultset []string) {
	output := make([][]string, 0)
	counter := 0

	for _, resultItem := range resultset {
		inner := make([]string, 0)
		inner = append(inner, strconv.Itoa(counter))
		inner = append(inner, resultItem)
		output = append(output, inner)
		counter++
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Count", "Article Title"})
	table.SetColMinWidth(1, 60)
	table.SetAutoWrapText(false)

	for _, v := range output {
		table.Append(v)
	}

	table.Render()
}
