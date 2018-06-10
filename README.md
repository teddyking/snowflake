# snowflake ❄️

A tool to help identify, triage and keep track of flaky tests.

## About

snowflake consists of four key components:

1. [reporter](reporter/) - a [ginkgo](http://onsi.github.io/ginkgo/) reporter, which sends the results of test runs to the snowflake server.
1. [server](cmd/snowflake/snowflake.go) - serves the snowflake API.
1. [web](cmd/snowflakeweb/snowflakeweb.go) - provides a web UI.
1. [snowgauge](snowgauge/) - the package responsible for actually detecting flaky tests.

## Installation and Usage

To use snowflake you must first install the snowflake server and web, which can be done as follows:

```
go get github.com/teddyking/snowflake
cd $GOPATH/src/github.com/teddyking/snowflake

make runserver
make runweb
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

Then finally run your test suites. Flaky tests will appear in the web UI, which by default is accessible at [http://localhost:8080](http://localhost:8080).

## Testing

[![CircleCI](https://circleci.com/gh/teddyking/snowflake/tree/master.svg?style=svg&circle-token=7d4e3f0b023e5aa60b8ce4461db19eb5472acad5)](https://circleci.com/gh/teddyking/snowflake/tree/master)

The tests can be run as follows:

```
make test            # run all tests
make testunit        # run unit tests only
make testintegration # run integration tests only
make teste2e         # run end-to-end tests only
```
