package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/teddyking/snowflake/web"
)

func main() {
	log.Println("--- snowflake web ---")

	listenPort := os.Getenv("PORT")
	if listenPort == "" {
		listenPort = "8080"
	}
	log.Printf("listening on port: %s", listenPort)

	fs := http.FileServer(http.Dir(filepath.Join("web", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", web.Index)

	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
