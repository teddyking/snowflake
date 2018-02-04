package main

import (
	"net/http"

	"github.com/teddyking/snowflake/server"
	"github.com/teddyking/snowflake/server/store"
)

func main() {
	inMemoryStore := store.NewInMemory()
	handler := server.NewHandler(inMemoryStore)

	http.ListenAndServe(":8080", handler)
}
