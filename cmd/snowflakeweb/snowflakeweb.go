package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/web/handler"
	"google.golang.org/grpc"
)

func main() {
	log.Println("--- snowflake web ---")

	serverPort := os.Getenv("SERVERPORT")
	if serverPort == "" {
		serverPort = "2929"
	}
	log.Printf("connecting to snowflake server on port: %s", serverPort)

	conn, err := grpc.Dial(":"+serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to server: %s", err.Error())
	}

	flakerService := api.NewFlakerClient(conn)

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "2930"
	}
	log.Printf("listening on port: %s", listenPort)

	staticDirPath := filepath.Join("web", "static")

	handler := handler.New(staticDirPath, flakerService)

	log.Fatal(http.ListenAndServe(":"+listenPort, handler))
}
