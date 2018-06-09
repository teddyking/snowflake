package snowgauge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/teddyking/snowflake/snowgauge"

	"github.com/teddyking/snowflake/api"
)

var _ = Describe("Snowgauge", func() {
	Describe("Flakes", func() {
		It("returns all flaky tests from the provided set of reports", func() {
			reports := []*api.Report{
				&api.Report{
					ImportPath:  "github.com/teddyking/snowflake",
					Commit:      "1",
					Description: "Integration Suite",
					Tests: []*api.Test{
						&api.Test{
							Description: "It is a flake",
							Location:    "/path/to/some_test.go:1",
							State:       api.Test_PASSED,
							StartedAt:   1528459728,
						},
						&api.Test{
							Description: "It is not a flake",
							Location:    "/path/to/some_test.go:2",
							State:       api.Test_PASSED,
							StartedAt:   1528459728,
						},
						&api.Test{
							Description: "It is still not a flake",
							Location:    "/path/to/some_test.go:3",
							State:       api.Test_FAILED,
							Failure:     &api.Failure{Message: "Not flaking just failing"},
							StartedAt:   1528459728,
						},
					},
				},
				&api.Report{
					ImportPath:  "github.com/teddyking/snowflake",
					Commit:      "1",
					Description: "Integration Suite",
					Tests: []*api.Test{
						&api.Test{
							Description: "It is a flake",
							Location:    "/path/to/some_test.go:1",
							State:       api.Test_FAILED,
							Failure:     &api.Failure{Message: "I flaked :("},
							StartedAt:   2528459728,
						},
						&api.Test{
							Description: "It is not a flake",
							Location:    "/path/to/some_test.go:2",
							State:       api.Test_PASSED,
							StartedAt:   2528459728,
						},
						&api.Test{
							Description: "It is still not a flake",
							Location:    "/path/to/some_test.go:3",
							State:       api.Test_FAILED,
							Failure:     &api.Failure{Message: "Not flaking just failing"},
							StartedAt:   2528459728,
						},
					},
				},
			}

			flakes, err := Flakes(reports)
			Expect(err).NotTo(HaveOccurred())

			Expect(len(flakes)).To(Equal(1))

			Expect(flakes[0].ImportPath).To(Equal("github.com/teddyking/snowflake"))
			Expect(flakes[0].Commit).To(Equal("1"))
			Expect(flakes[0].SuiteDescription).To(Equal("Integration Suite"))
			Expect(flakes[0].TestDescription).To(Equal("It is a flake"))
			Expect(flakes[0].Location).To(Equal("/path/to/some_test.go:1"))
			Expect(flakes[0].Failure.Message).To(Equal("I flaked :("))
			Expect(flakes[0].Successes).To(Equal(int64(1)))
			Expect(flakes[0].Failures).To(Equal(int64(1)))
			Expect(flakes[0].StartedAt).To(Equal(int64(1528459728)))
		})
	})
})
