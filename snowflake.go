package snowflake

import "time"

type Suite struct {
	Name  string
	Tests []*Test
}

type Test struct {
	Name        string
	CompletedAt time.Time
	State       string
	Failure     Failure
}

type Failure struct {
	Message string
}
