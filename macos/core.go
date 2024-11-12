package macos

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/elC0mpa/netstats/common"
)

func GetNetworkUsageByApp(searchTerm string) (map[string][2]float64, error) {
	output, err := runNettopCommand()
	if err != nil {
		return nil, err
	}

	appUsage := make(map[string][2]float64)
	scanner := bufio.NewScanner(&output)
	for scanner.Scan() {
		appName, sentMB, recvMB, err := parseNettopLine(scanner.Text())
		if err != nil || (sentMB <= 0.0 && recvMB <= 0.0) || (searchTerm != "" && !strings.Contains(strings.ToLower(appName), searchTerm)) {
			continue
		}
		common.AccumulateUsage(appUsage, appName, sentMB, recvMB)
	}
	return appUsage, scanner.Err()
}

func runNettopCommand() (bytes.Buffer, error) {
	cmd := exec.Command("nettop", "-P", "-L", "1", "-n", "-x")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	return output, err
}

func parseNettopLine(line string) (string, float64, float64, error) {
	fields := strings.Split(line, ",")
	if len(fields) < 6 {
		return "", 0, 0, fmt.Errorf("invalid line format")
	}

	appName := common.FormatAppName(fields[1])
	bytesSent, _ := strconv.ParseFloat(fields[5], 64)
	bytesRecv, _ := strconv.ParseFloat(fields[4], 64)

	return appName, bytesSent / 1024 / 1024, bytesRecv / 1024 / 1024, nil
}
