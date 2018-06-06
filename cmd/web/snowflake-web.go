package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/teddyking/snowflake/web/handler"
)

func main() {
	log.Println("--- snowflake web ---")

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "8080"
	}
	log.Printf("listening on port: %s", listenPort)

	templateDirPath := filepath.Join("web", "template")
	staticDirPath := filepath.Join("web", "static")

	handler := handler.New(templateDirPath, staticDirPath)

	log.Fatal(http.ListenAndServe(":"+listenPort, handler))
}
