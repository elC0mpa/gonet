package common

import (
	"github.com/elC0mpa/gonet/model/network"
)

func AccumulateUsage(appUsage map[string]network.NetworkInfo, appNetworkInfo network.AppNetworkInfo) {
	if usage, exists := appUsage[appNetworkInfo.AppName]; exists {
		appUsage[appNetworkInfo.AppName] = network.NetworkInfo{ReceivedBytes: usage.ReceivedBytes + appNetworkInfo.NetworkStats.ReceivedBytes, SentBytes: usage.SentBytes + appNetworkInfo.NetworkStats.SentBytes}
	} else {
		appUsage[appNetworkInfo.AppName] = network.NetworkInfo{ReceivedBytes: appNetworkInfo.NetworkStats.ReceivedBytes, SentBytes: appNetworkInfo.NetworkStats.SentBytes}
	}
}
