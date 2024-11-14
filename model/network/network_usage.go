package network

type NetworkUsage interface {
	GetNetworkUsageByApp(searchTerm string) (map[string][2]float64, error)
}
