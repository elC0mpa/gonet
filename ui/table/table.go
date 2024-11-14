package table

import (
	"fmt"
	"os"

	"github.com/elC0mpa/gonet/model/network"
	"github.com/jedib0t/go-pretty/v6/table"
)

type drawer struct {
	table      table.Writer
	rowsDrawed int
}

func NewTableDrawer() TableDrawer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Application", "Bytes Rcvd", "Bytes Sent"})

	return &drawer{
		table:      t,
		rowsDrawed: 0,
	}
}

func (tableDrawer *drawer) Draw(appUsage map[string]network.NetworkInfo) {
	if tableDrawer.rowsDrawed > 0 {
		tableDrawer.table.ResetRows()
	}

	for appName, netInfo := range appUsage {
		tableDrawer.table.AppendRow(table.Row{appName, netInfo.ReceivedBytes, netInfo.SentBytes})
	}

	fmt.Println(tableDrawer.table.Render())

	tableDrawer.rowsDrawed = len(appUsage)
}
