package table

import "github.com/elC0mpa/gonet/model/network"

type TableDrawer interface {
	Draw(appUsage map[string]network.NetworkInfo)
}
