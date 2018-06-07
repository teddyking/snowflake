package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/services/flaker"
	"github.com/teddyking/snowflake/services/reporter"
	"github.com/teddyking/snowflake/store"
	"google.golang.org/grpc"
)

func main() {
	log.Println("--- snowflake server ---")

	grpcServer := grpc.NewServer()
	store := store.NewVolatileStore()

	fakeFlakeAnalyser := func(reports []*api.Report) ([]*api.Flake, error) {
		return []*api.Flake{
			&api.Flake{
				ImportPath:       "github.com/teddyking/snowflake",
				Commit:           "cef64ee",
				SuiteDescription: "Integration Suite",
				TestDescription:  "It is a flake",
				Location:         "/some/path/to_test.go:12",
				Successes:        3,
				Failures:         3,
				StartedAt:        time.Now().Unix(),
				Failure:          &api.Failure{Message: "Expected 1 to equal 2"},
			},
		}, nil
	}

	reporterService := reporter.New(store)
	flakerService := flaker.New(store, fakeFlakeAnalyser)

	api.RegisterReporterServer(grpcServer, reporterService)
	api.RegisterFlakerServer(grpcServer, flakerService)

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		log.Fatal("PORT env var is not set")
	}
	log.Printf("listening on port: %s", listenPort)

	l, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		log.Fatalf("error listening on TCP port %s: %s", listenPort, err.Error())
	}

	log.Fatal(grpcServer.Serve(l))
}
