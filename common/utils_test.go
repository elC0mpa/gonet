package common_test

import (
	"testing"

	"github.com/elC0mpa/gonet/common"
	"github.com/elC0mpa/gonet/model/network"
	"github.com/stretchr/testify/assert"
)

func TestAccumulateUsage(t *testing.T) {
	appUsage := make(map[int]network.AppNetworkInfo)
	appNetworkInfo := network.AppNetworkInfo{
		Info: network.AppInfo{AppName: "Telegram", ProcessID: 9384},
		NetworkStats: network.NetworkInfo{
			ReceivedBytes: 100,
			SentBytes:     200,
		},
	}

	common.AccumulateUsage(appUsage, appNetworkInfo)

	assert.Equal(t, 100.0, appUsage[9384].NetworkStats.ReceivedBytes)
	assert.Equal(t, 200.0, appUsage[9384].NetworkStats.SentBytes)

	appNetworkInfo.NetworkStats.ReceivedBytes = 50
	appNetworkInfo.NetworkStats.SentBytes = 50
	common.AccumulateUsage(appUsage, appNetworkInfo)

	assert.Equal(t, 150.0, appUsage[9384].NetworkStats.ReceivedBytes)
	assert.Equal(t, 250.0, appUsage[9384].NetworkStats.SentBytes)
}
