package table

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/elC0mpa/gonet/model/network"
	"github.com/jedib0t/go-pretty/v6/table"
)

type drawer struct {
	table table.Writer
}

func NewTableDrawer() TableDrawer {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Application", "Bytes Rcvd", "Bytes Sent"})

	return &drawer{
		table: t,
	}
}

func clearConsole() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (tableDrawer *drawer) Draw(appUsage map[string]network.NetworkInfo) {
	tableDrawer.table.ResetRows()

	for appName, netInfo := range appUsage {
		tableDrawer.table.AppendRow(table.Row{appName, netInfo.ReceivedBytes, netInfo.SentBytes})
	}

	clearConsole()

	fmt.Println(tableDrawer.table.Render())
}
