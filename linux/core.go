//go:build linux
// +build linux

package linux

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func getNetworkUsageByApp(searchTerm string) (map[string][2]float64, error) {
	output, err := runSSCommand()
	if err != nil {
		return nil, err
	}

	appUsage := make(map[string][2]float64)
	scanner := bufio.NewScanner(&output)
	for scanner.Scan() {
		appName, sentMB, recvMB, err := parseSSLine(scanner.Text())
		if err != nil || (sentMB <= 0.0 && recvMB <= 0.0) || (searchTerm != "" && !strings.Contains(strings.ToLower(appName), searchTerm)) {
			continue
		}
		accumulateUsage(appUsage, appName, sentMB, recvMB)
	}
	return appUsage, scanner.Err()
}

func runSSCommand() (bytes.Buffer, error) {
	cmd := exec.Command("ss", "-tanp")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	return output, err
}

func parseSSLine(line string) (string, float64, float64, error) {
	fields := strings.Fields(line)
	if len(fields) < 6 || !strings.Contains(line, "pid=") {
		return "", 0, 0, fmt.Errorf("invalid line format")
	}

	var appName string
	for _, field := range fields {
		if strings.HasPrefix(field, "users:(") {
			appName = extractAppName(field)
			break
		}
	}

	sentMB := 0.01
	recvMB := 0.01

	return appName, sentMB, recvMB, nil
}

func extractAppName(field string) string {
	name := strings.Split(strings.Split(field, "\"")[1], ",")[0]
	words := strings.Fields(name)
	if len(words) > 2 {
		return strings.Join(words[:2], " ")
	}
	return name
}
