package macos

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/elC0mpa/gonet/common"
	"github.com/elC0mpa/gonet/model/network"
)

type MacNetworkUsage struct{}

func New() MacNetworkUsage {
	return MacNetworkUsage{}
}

func (ns MacNetworkUsage) GetNetworkUsageByApp(searchTerm string) (map[int]network.AppNetworkInfo, error) {
	output, err := ns.runCommand()
	if err != nil {
		return nil, err
	}

	appUsage := make(map[int]network.AppNetworkInfo)
	scanner := bufio.NewScanner(&output)
	for scanner.Scan() {
		netInfo, err := ns.parseCommand(scanner.Text())
		if err != nil || (netInfo.NetworkStats.SentBytes <= 0.0 && netInfo.NetworkStats.ReceivedBytes <= 0.0) || (searchTerm != "" && !strings.Contains(strings.ToLower(netInfo.Info.AppName), searchTerm)) {
			continue
		}
		common.AccumulateUsage(appUsage, netInfo)
	}
	return appUsage, scanner.Err()
}

func (ns MacNetworkUsage) runCommand() (bytes.Buffer, error) {
	cmd := exec.Command("nettop", "-P", "-L", "1", "-n", "-x")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	return output, err
}

func (ns MacNetworkUsage) parseCommand(line string) (network.AppNetworkInfo, error) {
	fields := strings.Split(line, ",")
	if len(fields) < 6 {
		return network.AppNetworkInfo{Info: network.AppInfo{}, NetworkStats: network.NetworkInfo{ReceivedBytes: 0, SentBytes: 0}}, fmt.Errorf("invalid line format")
	}

	bytesSent, err := strconv.ParseFloat(fields[5], 64)
	if err != nil {
		return network.AppNetworkInfo{}, fmt.Errorf("problem parsing bytes sent: %w", err)
	}

	bytesRecv, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		return network.AppNetworkInfo{}, fmt.Errorf("problem parsing bytes received: %w", err)
	}

	var networkInfo network.AppNetworkInfo = network.AppNetworkInfo{
		Info:         network.AppInfo{AppName: extractAppName(fields[1])},
		NetworkStats: network.NetworkInfo{ReceivedBytes: bytesRecv, SentBytes: bytesSent},
	}

	return networkInfo, nil
}

func extractAppName(field string) string {
	baseName := strings.Split(strings.TrimSpace(field), ".")[0]
	words := strings.Fields(baseName)
	if len(words) > 2 {
		return strings.Join(words[:2], " ")
	}
	return baseName
}
