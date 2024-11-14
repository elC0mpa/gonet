package table

import (
	"os"

	"github.com/elC0mpa/gonet/model/network"
	"github.com/jedib0t/go-pretty/v6/table"
)

func PrintUsageTable(appUsage map[string]network.NetworkInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"App Name", "Received Bytes", "Sent Bytes"})
	for appName, netInfo := range appUsage {
		t.AppendRow(table.Row{appName, netInfo.ReceivedBytes, netInfo.SentBytes})
		t.AppendSeparator()
	}
	t.Render()
}
