package store

import "github.com/teddyking/snowflake"

type inMemory struct {
	suites []snowflake.Suite
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (i *inMemory) All() ([]snowflake.Suite, error) {
	return i.suites, nil
}

func (i *inMemory) Save(suite snowflake.Suite) error {
	i.suites = append(i.suites, suite)
	return nil
}
