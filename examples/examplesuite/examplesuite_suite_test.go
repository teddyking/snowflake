package examplesuite_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/teddyking/snowflake"
	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func TestExamplesuite(t *testing.T) {
	RegisterFailHandler(Fail)

	importPath := "github.com/teddyking/snowflake/examples/examplesuite"
	gitCommitOutput, _ := exec.Command("git", "rev-parse", "HEAD").Output()
	gitCommit := strings.Trim(string(gitCommitOutput), "\n")

	conn, err := grpc.Dial(configureServerAddress(), configureDialOptions()...)
	if err != nil {
		log.Fatalf("could not connect to server: %s", err.Error())
	}

	reporterClient := api.NewReporterClient(conn)
	snowflakeReporter := snowflake.NewReporter(importPath, gitCommit, reporterClient)
	RunSpecsWithDefaultAndCustomReporters(t, "Examplesuite Suite", []Reporter{snowflakeReporter})
}

func configureServerAddress() string {
	serverHost := os.Getenv("SERVERHOST")
	if serverHost == "" {
		serverHost = "0.0.0.0"
	}
	serverPort := os.Getenv("SERVERPORT")
	if serverPort == "" {
		serverPort = "2929"
	}

	serverAddress := fmt.Sprintf("%s:%s", serverHost, serverPort)
	return serverAddress
}

func configureDialOptions() []grpc.DialOption {
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	tlsCrtPath := os.Getenv("TLSCRTPATH")
	if tlsCrtPath != "" {
		creds, err := credentials.NewClientTLSFromFile(tlsCrtPath, "")
		if err != nil {
			log.Fatalf("error reading TLS creds from '%s': %s", tlsCrtPath, err.Error())
		}
		creds.OverrideServerName(os.Getenv("SERVERHOST"))

		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	}

	return dialOpts
}
