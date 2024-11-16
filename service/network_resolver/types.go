package networkresolver

import "github.com/elC0mpa/gonet/model/network"

type NetworkResolver interface {
	GetNetworkUsage(OS string, searchTerm string) (map[int]network.AppNetworkInfo, error)
}
