package network

type NetworkUsage interface {
	GetNetworkUsageByApp(searchTerm string) (map[int]AppNetworkInfo, error)
}
