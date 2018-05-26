package examplesuite_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os/exec"
	"strings"
	"testing"

	"github.com/teddyking/snowflake"
)

func TestExamplesuite(t *testing.T) {
	RegisterFailHandler(Fail)

	codebase := "github.com/teddyking/snowflake/examples/examplesuite"
	gitCommitOutput, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	gitCommit := strings.Trim(string(gitCommitOutput), "\n")

	snowflakeReporter := snowflake.NewReporter(codebase, gitCommit)
	RunSpecsWithDefaultAndCustomReporters(t, "Examplesuite Suite", []Reporter{snowflakeReporter})
}
