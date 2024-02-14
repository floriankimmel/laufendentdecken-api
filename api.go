package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddress string
	store         Storage
}

func NewAPIServer(
	listenAddress string,
	store Storage,
) *APIServer {
	return &APIServer{
		listenAddress: listenAddress,
		store:         store,
	}
}

//nolint:all
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/reviews/{id}", makeHTTPHandleFunc(s.handleReviewsByID))
	router.HandleFunc("/trailEvents/{id}", makeHTTPHandleFunc(s.handleTrailEventsByID))

	log.Println("API server running on port: ", s.listenAddress)

	err := http.ListenAndServe(s.listenAddress, router)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if err := WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()}); err != nil {
				log.Println("Failed to write response: ", err)
			}

		}
	}
}
