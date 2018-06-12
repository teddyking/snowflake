package store

import (
	"sync"

	"github.com/teddyking/snowflake/api"
)

type VolatileStore struct {
	reports []*api.Report
	mu      sync.Mutex
}

type Opt func(*VolatileStore)

func WithInitialReports(reports []*api.Report) func(volatileStore *VolatileStore) {
	return func(volatileStore *VolatileStore) {
		volatileStore.reports = reports
	}
}

func NewVolatileStore(storeOpts ...Opt) *VolatileStore {
	volatileStore := &VolatileStore{}

	for _, storeOpt := range storeOpts {
		storeOpt(volatileStore)
	}

	return volatileStore
}

func (v *VolatileStore) CreateReport(report *api.Report) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.reports = append(v.reports, report)
	return nil
}

func (v *VolatileStore) ListReports() ([]*api.Report, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.reports, nil
}
