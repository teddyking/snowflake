package store

import (
	"sync"

	"github.com/teddyking/snowflake/api"
)

type VolatileStore struct {
	reports []*api.Report
	mu      sync.Mutex
}

func NewVolatileStore() *VolatileStore {
	return &VolatileStore{}
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
