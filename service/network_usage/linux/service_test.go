package linux

import (
	"testing"

	"github.com/elC0mpa/gonet/model/network"
	"github.com/stretchr/testify/assert"
)

func TestService_extractAppNameAndPid(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedName string
		expectedPid  int
	}{
		{name: "Extracting app name and pid linux for chrome", input: `users:(("chrome",pid=8220,fd=24))`, expectedName: "chrome", expectedPid: 8220},
		{name: "Extracting app name and pid linux for protonvpn", input: `users:(("protonvpn-app",pid=7156,fd=37))`, expectedName: "protonvpn-app", expectedPid: 7156},
		{name: "Extracting app name and pid linux for kitty", input: `users:(("kitty",pid=772639,fd=13))`, expectedName: "kitty", expectedPid: 772639},
	}

	service := &LinuxNetworkUsage{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pId, appName := service.extractAppNameAndPID(test.input)
			assert.Equal(t, test.expectedName, appName)
			assert.Equal(t, test.expectedPid, pId)
		})
	}
}

func TestService_parseAppNameAndId(t *testing.T) {
	service := &LinuxNetworkUsage{}

	appUsage := make(map[int]network.AppNetworkInfo)
	line := `tcp          ESTAB             0             0                                                                                                  10.2.0.2:58390                           107.23.218.0:443             users:(("chrome",pid=8220,fd=27))`

	service.parseAppNameAndId(line, appUsage)

	appNetworkInfo, exists := appUsage[8220]

	assert.Equal(t, true, exists)
	assert.Equal(t, 1, len(appUsage))

	service.parseAppNameAndId(line, appUsage)
	assert.Equal(t, 1, len(appUsage))

	assert.Equal(t, 8220, appNetworkInfo.Info.ProcessID)
	assert.Equal(t, "chrome", appNetworkInfo.Info.AppName)
}

func TestService_sumNetworkBytes(t *testing.T) {
	service := &LinuxNetworkUsage{}

	input := `Inter-|   Receive                                                |  Transmit
 face |bytes    packets errs drop fifo frame compressed multicast|bytes    packets errs drop fifo colls carrier compressed
    lo: 215465903  298447    0    0    0     0          0         0 215465903  298447    0    0    0     0       0          0
  wlo1: 9275431231 8074380    0    0    0     0          0         0 1894672559 6261018    0  559    0     0       0          0
br-bbc8e71bf5a3:       0       0    0    0    0     0          0         0        0       0    0  503    0     0       0          0
br-10fa63e1ba5a:       0       0    0    0    0     0          0         0        0       0    0  504    0     0       0          0
br-31ace59d3fea:       0       0    0    0    0     0          0         0        0       0    0  504    0     0       0          0
docker0: 3949813   12952    0    0    0     0          0         0  4898952   28628    0   44    0     0       0          0
veth1a5b657: 4131141   12952    0    0    0     0          0         0  5180101   29142    0    0    0     0       0          0
pvpnksintrf0:       0       0    0    0    0     0          0         0    32994     133    0    0    0     0       0          0
proton0: 3995644   12949    0    0    0     0          0         0  3640200    8473    0    0    0     0       0          0`

	rcvdBytes, sentBytes := service.sumNetworkBytes(input)
	assert.Equal(t, 9502973732, rcvdBytes)
	assert.Equal(t, 2123890709, sentBytes)
}
