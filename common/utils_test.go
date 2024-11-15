package common_test

import (
	"testing"

	"github.com/elC0mpa/gonet/common"
	"github.com/elC0mpa/gonet/model/network"
	"github.com/stretchr/testify/assert"
)

// func TestFormatAppName(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected string
// 	}{
// 		{input: "12:31:48.335784,launchd.1,,,0,0,0,0,0,,,,,,,,,,,,", expected: "launchd"},
// 		{input: "12:31:48.335833,Telegram.38574,,,808446,164427,0,0,0,,,,,,,,,,,,", expected: "Telegram"},
// 	}
//
// 	for _, test := range tests {
// 		t.Run(test.input, func(t *testing.T) {
// 			result := common.FormatAppName(test.input)
// 			assert.Equal(t, test.expected, result)
// 		})
// 	}
// }

func TestAccumulateUsage(t *testing.T) {
	appUsage := make(map[string]network.NetworkInfo)
	appNetworkInfo := network.AppNetworkInfo{
		AppName: "Telegram",
		NetworkStats: network.NetworkInfo{
			ReceivedBytes: 100,
			SentBytes:     200,
		},
	}

	common.AccumulateUsage(appUsage, appNetworkInfo)

	assert.Equal(t, 100.0, appUsage["Telegram"].ReceivedBytes)
	assert.Equal(t, 200.0, appUsage["Telegram"].SentBytes)

	appNetworkInfo.NetworkStats.ReceivedBytes = 50
	appNetworkInfo.NetworkStats.SentBytes = 50
	common.AccumulateUsage(appUsage, appNetworkInfo)

	assert.Equal(t, 150.0, appUsage["Telegram"].ReceivedBytes)
	assert.Equal(t, 250.0, appUsage["Telegram"].SentBytes)
}
