package macos

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/elC0mpa/netstats/common"
	"github.com/elC0mpa/netstats/model/network"
)

type MacNetworkUsage struct{}

func New() MacNetworkUsage {
	return MacNetworkUsage{}
}

func (ns MacNetworkUsage) GetNetworkUsageByApp(searchTerm string) (map[string][2]float64, error) {
	output, err := ns.runCommand()
	if err != nil {
		return nil, err
	}

	appUsage := make(map[string][2]float64)
	scanner := bufio.NewScanner(&output)
	for scanner.Scan() {
		netInfo, err := ns.parseCommand(scanner.Text())
		if err != nil || (netInfo.SentBytes <= 0.0 && netInfo.ReceivedBytes <= 0.0) || (searchTerm != "" && !strings.Contains(strings.ToLower(netInfo.AppName), searchTerm)) {
			continue
		}
		common.AccumulateUsage(appUsage, netInfo.AppName, netInfo.SentBytes, netInfo.ReceivedBytes)
	}
	return appUsage, scanner.Err()
}

func (ns MacNetworkUsage) runCommand() (bytes.Buffer, error) {
	cmd := exec.Command("nettop", "-P", "-L", "1", "-n", "-x")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	return output, err
}

func (ns MacNetworkUsage) parseCommand(line string) (network.NetworkInfo, error) {
	fields := strings.Split(line, ",")
	if len(fields) < 6 {
		return network.NetworkInfo{AppName: "", ReceivedBytes: 0, SentBytes: 0}, fmt.Errorf("invalid line format")
	}

	bytesSent, err := strconv.ParseFloat(fields[5], 64)
	if err != nil {
		panic("Problem parsing bytes sent")
	}

	bytesRecv, err := strconv.ParseFloat(fields[4], 64)
	if err != nil {
		panic("Problem parsing bytes received")
	}

	var networkInfo network.NetworkInfo = network.NetworkInfo{
		AppName:       common.FormatAppName(fields[1]),
		ReceivedBytes: bytesRecv,
		SentBytes:     bytesSent,
	}

	return networkInfo, nil
}
