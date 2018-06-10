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
)

func main() {
	log.Println("--- snowflake server ---")

	grpcServer := grpc.NewServer()
	store := store.NewVolatileStore()

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
