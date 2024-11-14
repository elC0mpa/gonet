package network

type NetworkUsage interface {
	GetNetworkUsageByApp(searchTerm string) (map[string]NetworkInfo, error)
}
