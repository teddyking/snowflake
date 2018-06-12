package main

import (
	"log"
	"net"
	"os"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/services/flaker"
	"github.com/teddyking/snowflake/services/reporter"
	"github.com/teddyking/snowflake/snowgauge"
	"github.com/teddyking/snowflake/store"
	"google.golang.org/grpc"

	testdata "github.com/teddyking/snowflake/test/data"
)

func main() {
	log.Println("--- snowflake server ---")

	grpcServer := grpc.NewServer()

	storeOpts := []store.Opt{}
	if os.Getenv("SEEDSTORE") == "true" {
		storeOpts = append(storeOpts, store.WithInitialReports(testdata.ReportsWithAFlake))
	}
	store := store.NewVolatileStore(storeOpts...)

	reporterService := reporter.New(store)
	flakerService := flaker.New(store, snowgauge.Flakes)

	api.RegisterReporterServer(grpcServer, reporterService)
	api.RegisterFlakerServer(grpcServer, flakerService)

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "2929"
	}
	log.Printf("listening on port: %s", listenPort)

	l, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		log.Fatalf("error listening on TCP port %s: %s", listenPort, err.Error())
	}

	log.Fatal(grpcServer.Serve(l))
}
