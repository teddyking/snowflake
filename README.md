# snowflake ❄️

A tool to help identify, triage and keep track of flaky tests.

## About

snowflake consists of four key components:

1. Reporter - a [Ginkgo](http://onsi.github.io/ginkgo/) reporter, which sends the results of test runs to the snowflake server.
1. Server - serves the snowflake API.
1. Web - provides a web UI.
1. snowgauge - the package responsible for actually detecting flaky tests.

## Installation and Usage

To use snowflake you must first install the snowflake server and web, which can be done as follows:

```
# install binaries
go get github.com/teddyking/snowflake/cmd/snowflake
go get github.com/teddyking/snowflake/cmd/snowflakeweb

# run them
PORT=2929 snowflake
SERVERPORT=2929 PORT=8080 snowflakeweb
```

Then you need to configure your ginkgo test suites with the snowflake reporter, for example:

```
# example_suite_test.go

func TestExamplesuite(t *testing.T) {
	RegisterFailHandler(Fail)

	importPath := "github.com/teddyking/snowflake"
	gitCommitOutput, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	gitCommit := strings.Trim(string(gitCommitOutput), "\n")

	conn, err := grpc.Dial(":2929", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %s", err.Error())
	}
	reporterClient := api.NewReporterClient(conn)

	snowflakeReporter := snowflake.NewReporter(importPath, gitCommit, reporterClient)

	RunSpecsWithDefaultAndCustomReporters(t, "Examplesuite Suite", []Reporter{snowflakeReporter})
}
```

For a full example see the [examples/examplesuite](examples/examplesuite).

Then finally run your test suites. Flaky tests will appear in the web UI.

## Running snowflake's tests

`make test`
