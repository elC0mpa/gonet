package table

import (
	"fmt"

	"github.com/elC0mpa/gonet/common"
	"github.com/elC0mpa/gonet/model/network"
	"github.com/jedib0t/go-pretty/v6/table"
)

type drawer struct {
	table table.Writer
}

func NewTableDrawer() TableDrawer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Pid", "Application", "Bytes Rcvd", "Bytes Sent"})
	return &drawer{
		table: t,
	}
}

func (tableDrawer *drawer) Draw(appUsage map[int]network.AppNetworkInfo) {
	tableDrawer.table.ResetRows()

	for pId, appNetInfo := range appUsage {
		tableDrawer.table.AppendRow(table.Row{pId, appNetInfo.Info.AppName, common.FromBytesToString(appNetInfo.NetworkStats.ReceivedBytes), common.FromBytesToString(appNetInfo.NetworkStats.SentBytes)})
	}

	// common.ClearConsole()

	fmt.Println(tableDrawer.table.Render())
}
