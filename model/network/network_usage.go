package network

import "bytes"

type NetworkUsage interface {
	GetNetworkUsageByApp(searchTerm string) (map[string][2]float64, error)
	runCommand() (bytes.Buffer, error)
	parseCommand(line string) (NetworkInfo, error)
}
