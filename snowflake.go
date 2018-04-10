package snowflake

import "time"

type Suite struct {
	Name        string    `json:"name"`
	Commit      string    `json:"commit"`
	Tests       []*Test   `json:"tests"`
	StartedAt   time.Time `json:"startedAt"`
	CompletedAt time.Time `json:"completedAt"`
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
