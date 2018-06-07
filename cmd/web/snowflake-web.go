package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/web/handler"
	"github.com/teddyking/snowflake/web/handler/handlerfakes"
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
	flakerService := new(handlerfakes.FakeFlakerService)

	flakerService.ListReturns(&api.FlakerListRes{
		Flakes: []*api.Flake{
			&api.Flake{
				ImportPath:       "github.com/teddyking/snowflake",
				Commit:           "3a76bbc",
				SuiteDescription: "Integration Suite",
				TestDescription:  "[It] does something successfully",
				Location:         "/path/to/some_test.go:12",
				Successes:        10,
				Failures:         3,
				StartedAt:        time.Now().Unix(),
				Failure:          &api.Failure{Message: "Expected 1 to equal 2"},
			},
		},
	}, nil)

	handler := handler.New(templateDirPath, staticDirPath, flakerService)

	log.Fatal(http.ListenAndServe(":"+listenPort, handler))
}
