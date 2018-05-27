package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/teddyking/snowflake/api"
	"google.golang.org/grpc"
)

func main() {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		log.Fatal("PORT env var is not set")
	}
	log.Printf("connecting to port: %s", serverPort)

	conn, err := grpc.Dial(":"+serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %s", err.Error())
	}

	client := api.NewSuiteClient(conn)

	ctx := context.Background()
	req := &api.ListRequest{}
	res, err := client.List(ctx, req)
	if err != nil {
		log.Fatalf("error listing summaries: %s", err.Error())
	}

	for _, summary := range res.SuiteSummaries {
		printSummary(summary)
	}

}

func printSummary(summary *api.SuiteSummary) {
	fmt.Printf("Codebase: %s\n", summary.Codebase)
	fmt.Printf("Commit: %s\n", summary.Commit)
	fmt.Printf("Name: %s\n", summary.Name)

	for _, test := range summary.Tests {
		fmt.Printf("\t%s: %s\n", test.State.String(), test.Name)
	}
}
