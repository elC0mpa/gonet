package common

import (
	"fmt"
	"strings"
)

func FormatAppName(fullName string) string {
	baseName := strings.Split(strings.TrimSpace(fullName), ".")[0]
	words := strings.Fields(baseName)
	if len(words) > 2 {
		return strings.Join(words[:2], " ")
	}
	return baseName
}

func AccumulateUsage(appUsage map[string][2]float64, appName string, sentMB, recvMB float64) {
	if usage, exists := appUsage[appName]; exists {
		appUsage[appName] = [2]float64{usage[0] + sentMB, usage[1] + recvMB}
	} else {
		appUsage[appName] = [2]float64{sentMB, recvMB}
	}
}

func PrintUsageTable(appUsage map[string][2]float64) {
	fmt.Printf("\033[1;34m%-30s %-15s %-15s\033[0m\n", "Application", "Sent", "Received")

	rowColor := "\033[0m"
	rowIndex := 0
	for app, usage := range appUsage {
		if rowIndex%2 == 0 {
			rowColor = "\033[38;5;242m"
		} else {
			rowColor = "\033[0m"
		}
		sentStr := fmt.Sprintf("%.2f MB", usage[0])
		recvStr := fmt.Sprintf("%.2f MB", usage[1])
		fmt.Printf("%s%-30s %-15s %-15s\033[0m\n", rowColor, app, sentStr, recvStr)
		rowIndex++
	}
}
