package examplesuite_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"log"
	"os/exec"
	"strings"
	"testing"

	"github.com/teddyking/snowflake"
	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
)

func TestExamplesuite(t *testing.T) {
	RegisterFailHandler(Fail)

	importPath := "github.com/teddyking/snowflake/examples/examplesuite"
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
