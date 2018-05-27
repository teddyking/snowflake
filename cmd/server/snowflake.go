package main

import (
	"log"
	"net"
	"os"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/server"
	"github.com/teddyking/snowflake/store"
	"google.golang.org/grpc"
)

func main() {
	log.Println("--- snowflake server ---")

	grpcServer := grpc.NewServer()
	suiteServer := server.New(store.NewVolatileStore())
	api.RegisterSuiteServer(grpcServer, suiteServer)

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
