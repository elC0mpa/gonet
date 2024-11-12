package linux

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/elC0mpa/netstats/common"
	"github.com/elC0mpa/netstats/model/network"
)

func GetNetworkUsageByApp(searchTerm string) (map[string][2]float64, error) {
	output, err := runCommand()
	if err != nil {
		return nil, err
	}

	appUsage := make(map[string][2]float64)
	scanner := bufio.NewScanner(&output)
	for scanner.Scan() {
		networkInfo, err := parseCommand(scanner.Text())
		if err != nil || (networkInfo.SentBytes <= 0.0 && networkInfo.ReceivedBytes <= 0.0) || (searchTerm != "" && !strings.Contains(strings.ToLower(networkInfo.AppName), searchTerm)) {
			continue
		}
		common.AccumulateUsage(appUsage, networkInfo.AppName, networkInfo.SentBytes, networkInfo.ReceivedBytes)
	}
	return appUsage, scanner.Err()
}

func runCommand() (bytes.Buffer, error) {
	cmd := exec.Command("ss", "-tanp")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	return output, err
}

func parseCommand(line string) (network.NetworkInfo, error) {
	fields := strings.Fields(line)
	if len(fields) < 6 || !strings.Contains(line, "pid=") {
		return network.NetworkInfo{AppName: "", ReceivedBytes: 0, SentBytes: 0}, fmt.Errorf("invalid line format")
	}

	var networkInfo network.NetworkInfo = network.NetworkInfo{SentBytes: 0.01, ReceivedBytes: 0.01}
	for _, field := range fields {
		if strings.HasPrefix(field, "users:(") {
			networkInfo.AppName = extractAppName(field)
			break
		}
	}

	return networkInfo, nil
}

func extractAppName(field string) string {
	name := strings.Split(strings.Split(field, "\"")[1], ",")[0]
	words := strings.Fields(name)
	if len(words) > 2 {
		return strings.Join(words[:2], " ")
	}
	return name
}
