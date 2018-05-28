package store

import (
	"sync"

	"github.com/teddyking/snowflake/api"
)

type VolatileStore struct {
	summaries []*api.SuiteSummary
	mu        sync.Mutex
}

func NewVolatileStore() *VolatileStore {
	return &VolatileStore{}
}

func (v *VolatileStore) Create(suiteSummary *api.SuiteSummary) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.summaries = append(v.summaries, suiteSummary)
	return nil
}

func (v *VolatileStore) List() ([]*api.SuiteSummary, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	return v.summaries, nil
}