package main

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/services/flaker"
	"github.com/teddyking/snowflake/services/reporter"
	"github.com/teddyking/snowflake/snowgauge"
	"github.com/teddyking/snowflake/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	testdata "github.com/teddyking/snowflake/test/data"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	log.Printf("starting snowflake")

	grpcServer := grpc.NewServer(configureServerOptions()...)
	store := store.NewVolatileStore(configureStoreOptions()...)

	reporterService := reporter.New(store)
	flakerService := flaker.New(store, snowgauge.Flakes)

	api.RegisterReporterServer(grpcServer, reporterService)
	api.RegisterFlakerServer(grpcServer, flakerService)

	listenAddress := configureListenAddress()
	l, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("error listening on TCP address %s: %s", listenAddress, err.Error())
	}

	log.Fatal(grpcServer.Serve(l))
}

func configureServerOptions() []grpc.ServerOption {
	serverOpts := []grpc.ServerOption{}

	tlsKeyPath := os.Getenv("TLSKEYPATH")
	tlsCrtPath := os.Getenv("TLSCRTPATH")
	if tlsKeyPath != "" && tlsCrtPath != "" {
		creds, err := credentials.NewServerTLSFromFile(tlsCrtPath, tlsKeyPath)
		if err != nil {
			log.Fatalf("error reading TLS creds from '%s', '%s': %s", tlsCrtPath, tlsKeyPath, err.Error())
		}

		serverOpts = append(serverOpts, grpc.Creds(creds))
		log.Printf("tls configured")
	}

	return serverOpts
}

func configureStoreOptions() []store.Opt {
	storeOpts := []store.Opt{}

	if os.Getenv("SEEDSTORE") == "true" {
		storeOpts = append(storeOpts, store.WithInitialReports(testdata.ReportsWithAFlake))
	}

	return storeOpts
}

func configureListenAddress() string {
	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "2929"
	}

	return fmt.Sprintf("localhost:%s", listenPort)
}
