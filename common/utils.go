package common

import (
	"strings"

	"github.com/elC0mpa/gonet/model/network"
)

func FormatAppName(fullName string) string {
	baseName := strings.Split(strings.TrimSpace(fullName), ".")[0]
	words := strings.Fields(baseName)
	if len(words) > 2 {
		return strings.Join(words[:2], " ")
	}
	return baseName
}

func AccumulateUsage(appUsage map[string]network.NetworkInfo, appNetworkInfo network.AppNetworkInfo) {
	if usage, exists := appUsage[appNetworkInfo.AppName]; exists {
		appUsage[appNetworkInfo.AppName] = network.NetworkInfo{ReceivedBytes: usage.ReceivedBytes + appNetworkInfo.NetworkStats.ReceivedBytes, SentBytes: usage.SentBytes + appNetworkInfo.NetworkStats.SentBytes}
	} else {
		appUsage[appNetworkInfo.AppName] = network.NetworkInfo{ReceivedBytes: appNetworkInfo.NetworkStats.ReceivedBytes, SentBytes: appNetworkInfo.NetworkStats.SentBytes}
	}
}
