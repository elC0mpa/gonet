package common

import (
	"github.com/elC0mpa/gonet/model/network"
)

func AccumulateUsage(appUsage map[int]network.AppNetworkInfo, appNetworkInfo network.AppNetworkInfo) {
	if usage, exists := appUsage[appNetworkInfo.Info.ProcessID]; exists {
		appUsage[appNetworkInfo.Info.ProcessID] = network.AppNetworkInfo{NetworkStats: network.NetworkInfo{ReceivedBytes: usage.NetworkStats.ReceivedBytes + appNetworkInfo.NetworkStats.ReceivedBytes, SentBytes: usage.NetworkStats.SentBytes + appNetworkInfo.NetworkStats.SentBytes}, Info: network.AppInfo{ProcessID: usage.Info.ProcessID, AppName: usage.Info.AppName}}
	} else {
		appUsage[appNetworkInfo.Info.ProcessID] = network.AppNetworkInfo{NetworkStats: network.NetworkInfo{ReceivedBytes: appNetworkInfo.NetworkStats.ReceivedBytes, SentBytes: appNetworkInfo.NetworkStats.SentBytes}, Info: network.AppInfo{ProcessID: appNetworkInfo.Info.ProcessID, AppName: appNetworkInfo.Info.AppName}}
	}
}
