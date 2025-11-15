package output

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func NewTable(headers []string) *tablewriter.Table {
	n := len(headers)
	colors := make([]tablewriter.Colors, n)
	for i := range colors {
		colors[i] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetBorder(true)
	table.SetRowLine(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderColor(colors...)
	return table
}
