package networkresolver

import (
	"fmt"

	"github.com/elC0mpa/gonet/model/network"
)

type resolver struct {
	clients map[string]network.NetworkUsage
}

func NewNetworkResolver(clients map[string]network.NetworkUsage) NetworkResolver {
	return resolver{
		clients: clients,
	}
}

func (nr resolver) GetNetworkUsage(OS string, searchTerm string) (map[string]network.NetworkInfo, error) {
	client, ok := nr.clients[OS]

	if !ok {
		return nil, fmt.Errorf("Couldn't resolve network usage for OS %s", OS)
	}

	networkUsage, err := client.GetNetworkUsageByApp(searchTerm)
	if err != nil {
		return nil, err
	}

	return networkUsage, nil
}
