package linux

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/elC0mpa/gonet/model/network"
)

type LinuxNetworkUsage struct{}

func New() LinuxNetworkUsage {
	return LinuxNetworkUsage{}
}

func (ns LinuxNetworkUsage) GetNetworkUsageByApp(searchTerm string) (map[int]network.AppNetworkInfo, error) {
	output, err := ns.runCommand()
	if err != nil {
		return nil, err
	}

	appUsage := make(map[int]network.AppNetworkInfo)
	scanner := bufio.NewScanner(&output)
	for scanner.Scan() {
		pid, processIdAlreadyProcessed, err := ns.parseAppNameAndId(scanner.Text(), appUsage)
		if err == nil && processIdAlreadyProcessed == false {
			filePath := fmt.Sprintf("/proc/%d/net/dev", pid)

			networkFileForProcess, err := os.Open(filePath)
			if err != nil {
				panic("Can't open file for process")
			}

			content, err := io.ReadAll(networkFileForProcess)
			networkFileForProcess.Close()
			if err != nil {
				panic("Can't read file content")
			}

			contentStr := string(content)

			fmt.Println("Data read from file: ", filePath, "\n", contentStr)
			rcvdBytes, sentBytes := ns.sumNetworkBytes(contentStr)

			appNetworkInfo := appUsage[pid]
			appUsage[pid] = network.AppNetworkInfo{Info: network.AppInfo{ProcessID: appNetworkInfo.Info.ProcessID, AppName: appNetworkInfo.Info.AppName}, NetworkStats: network.NetworkInfo{ReceivedBytes: float64(rcvdBytes), SentBytes: float64(sentBytes)}}
		}
	}

	return appUsage, scanner.Err()
}

func (ns LinuxNetworkUsage) sumNetworkBytes(data string) (int, int) {
	lines := strings.Split(data, "\n")
	var totalReceived, totalTransmitted int

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 10 || strings.HasSuffix(parts[0], "|") {
			continue
		}

		receivedBytes, err1 := strconv.Atoi(parts[1])
		transmittedBytes, err2 := strconv.Atoi(parts[9])
		if err1 == nil && err2 == nil {
			totalReceived += receivedBytes
			totalTransmitted += transmittedBytes
		}
	}

	return totalReceived, totalTransmitted
}

func (ns LinuxNetworkUsage) runCommand() (bytes.Buffer, error) {
	cmd := exec.Command("bash", "-c", "ss -utanp | awk '/ESTAB/ && /pid=/ {print $0}' ")
	var output bytes.Buffer
	cmd.Stdout = &output
	err := cmd.Run()
	return output, err
}

func (ns LinuxNetworkUsage) parseAppNameAndId(line string, appUsage map[int]network.AppNetworkInfo) (int, bool, error) {
	fields := strings.Fields(line)
	if len(fields) < 6 || !strings.Contains(line, "pid=") {
		return 0, false, fmt.Errorf("invalid line format")
	}

	var pid int
	var name string
	var exists bool

	for _, field := range fields {
		if strings.HasPrefix(field, "users:(") {
			pid, name = ns.extractAppNameAndPID(field)
			if _, exists = appUsage[pid]; exists == false && name != "" {
				appUsage[pid] = network.AppNetworkInfo{Info: network.AppInfo{ProcessID: pid, AppName: name}}
			}
		}
	}

	return pid, exists, nil
}

func (ns LinuxNetworkUsage) extractAppNameAndPID(field string) (int, string) {
	name := strings.Split(strings.Split(field, "\"")[1], ",")[0]
	words := strings.Fields(name)
	appName := name
	if len(words) > 2 {
		appName = strings.Join(words[:2], " ")
	}

	pidStr := strings.Split(strings.Split(field, "pid=")[1], ",")[0]
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, ""
	}

	return pid, appName
}
