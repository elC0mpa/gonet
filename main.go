package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/elC0mpa/gonet/model/network"
	networkresolver "github.com/elC0mpa/gonet/service/network_resolver"
	"github.com/elC0mpa/gonet/service/network_usage/linux"
	"github.com/elC0mpa/gonet/service/network_usage/macos"
	"github.com/elC0mpa/gonet/ui/table"
)

func main() {
	printOSInfo()

	var searchTerm string
	if len(os.Args) > 1 {
		searchTerm = strings.ToLower(os.Args[1])
	}

	linuxClient := linux.New()
	macClient := macos.New()

	resolver := networkresolver.NewNetworkResolver(
		map[string]network.NetworkUsage{
			"linux":  linuxClient,
			"darwin": macClient,
		})

	appUsage, err := resolver.GetNetworkUsage(runtime.GOOS, searchTerm)
	if err != nil {
		panic(err)
	}
	
	table.PrintUsageTable(appUsage)
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
