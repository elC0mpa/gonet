package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/elC0mpa/netstats/common"
	"github.com/elC0mpa/netstats/linux"
	"github.com/elC0mpa/netstats/macos"
)

func main() {
	printOSInfo()

	var searchTerm string
	if len(os.Args) > 1 {
		searchTerm = strings.ToLower(os.Args[1])
	}

	var appUsage map[string][2]float64
	var err error

	switch runtime.GOOS {
	case "darwin":
		appUsage, err = macos.GetNetworkUsageByApp(searchTerm)
	case "linux":
		appUsage, err = linux.GetNetworkUsageByApp(searchTerm)
	}

	if err != nil {
		panic(err)
	}

	common.PrintUsageTable(appUsage)
}

func printOSInfo() {
	switch runtime.GOOS {
	case "darwin":
		fmt.Printf("Operating System: macOS\n")
	case "linux":
		fmt.Printf("Operating System: Linux\n")
	case "windows":
		fmt.Printf("Operating System: Windows\n")
	default:
		fmt.Printf("Operating System: %s\n", runtime.GOOS)
	}
}
