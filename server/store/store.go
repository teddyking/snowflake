package store

import "github.com/teddyking/snowflake"

//go:generate counterfeiter . Store

type Store interface {
	All() ([]snowflake.Suite, error)
	Save(suite snowflake.Suite) error
}
