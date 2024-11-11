package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	printOSInfo()

	var searchTerm string
	if len(os.Args) > 1 {
		searchTerm = strings.ToLower(os.Args[1])
	}

	appUsage, err := getNetworkUsageByApp(searchTerm)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	printUsageTable(appUsage)
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
