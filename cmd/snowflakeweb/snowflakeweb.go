package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/web/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.Println("--- snowflake web ---")

	conn, err := grpc.Dial(configureServerAddress(), configureDialOptions()...)
	if err != nil {
		log.Fatalf("could not connect to server: %s", err.Error())
	}

	flakerService := api.NewFlakerClient(conn)

	handler := handler.New(configureStaticDirPath(), flakerService)

	log.Fatal(http.ListenAndServe(configureListenAddress(), handler))
}

func configureServerAddress() string {
	serverPort := os.Getenv("SERVERPORT")
	if serverPort == "" {
		serverPort = "2929"
	}

	serverAddress := fmt.Sprintf("localhost:%s", serverPort)
	log.Printf("server address set to: %s", serverAddress)

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

		dialOpts = []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	}

	return dialOpts
}

func configureListenAddress() string {
	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "2930"
	}

	listenAddress := fmt.Sprintf("localhost:%s", listenPort)
	log.Printf("listen address set to: %s", listenAddress)

	return listenAddress
}

func configureStaticDirPath() string {
	staticDirPath := os.Getenv("STATICDIR")
	if staticDirPath == "" {
		staticDirPath = filepath.Join("web", "static")
	}
	log.Printf("serving static assets from: %s", staticDirPath)

	return staticDirPath
}
