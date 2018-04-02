package snowflake

import "time"

type Suite struct {
	Name   string  `json:"name"`
	Commit string  `json:"commit"`
	Tests  []*Test `json:"tests"`
}

type Test struct {
	Name        string    `json:"name"`
	CompletedAt time.Time `json:"completedAt"`
	State       string    `json:"state"`
	Failure     Failure   `json:"failure"`
}

type Failure struct {
	Message string `json:"message"`
}
