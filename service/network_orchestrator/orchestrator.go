package networkorchestrator

import (
	"fmt"
	"runtime"
	"time"

	networkresolver "github.com/elC0mpa/gonet/service/network_resolver"
	"github.com/elC0mpa/gonet/ui/table"
)

type orchestrator struct {
	resolver   networkresolver.NetworkResolver
	timer      *time.Ticker
	searchTerm string
	table      table.TableDrawer
}

func NewNetworkOrchestrator(resolver networkresolver.NetworkResolver, drawer table.TableDrawer, searchTerm string) NetworkOrchestrator {
	timer := time.NewTicker(1 * time.Second)
	timer.Stop()

	return orchestrator{
		resolver:   resolver,
		timer:      timer,
		table:      drawer,
		searchTerm: searchTerm,
	}
}

func (orchestrator orchestrator) Start() chan bool {
	done := make(chan bool, 1)
	orchestrator.timer.Reset(1 * time.Second)

	go func() {
		defer orchestrator.Stop()
		for {
			select {
			case <-orchestrator.timer.C:
				appUsage, err := orchestrator.resolver.GetNetworkUsage(runtime.GOOS, orchestrator.searchTerm)
				if err != nil {
					orchestrator.timer.Stop()
					panic(fmt.Errorf("Error when gathering network usage: %s", err))
				}

				orchestrator.table.Draw(appUsage)

			case <-done:
				return
			}
		}
	}()

	return done
}

func (orchestrator orchestrator) Stop() {
	orchestrator.timer.Stop()
}
