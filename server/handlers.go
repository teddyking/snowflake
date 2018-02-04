package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teddyking/snowflake"
	"github.com/teddyking/snowflake/server/store"
)

func NewHandler(store store.Store) http.Handler {
	r := mux.NewRouter()

	suitesHandler := &SuitesHandler{store}

	r.HandleFunc("/v1/suites", suitesHandler.List).Methods("GET")
	r.HandleFunc("/v1/suites", suitesHandler.Create).Methods("POST")

	return r
}

type SuitesHandler struct {
	store store.Store
}

func NewSuitesHandler(store store.Store) *SuitesHandler {
	return &SuitesHandler{
		store: store,
	}
}

func (h *SuitesHandler) List(w http.ResponseWriter, req *http.Request) {
	suites, err := h.store.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(suites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
	return
}

func (h *SuitesHandler) Create(w http.ResponseWriter, req *http.Request) {
	var suite snowflake.Suite

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &suite); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.store.Save(suite); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
