package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// WriteJSON writes the JSON representation of v to the response.
func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// apiFunc is a handler that returns an error.
type apiFunc func(w http.ResponseWriter, r *http.Request) error

// apiError is an error that can be returned by an apiFunc.
type apiError struct {
	Error string
}

// makeHTTPHandler creates an http.HandlerFunc from an apiFunc.
func makeHTTPHandlerFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err == nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

// APIServer is a server for the API.

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

// start server
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handlerAccount))

	log.Println(("JSON API server listening on " + s.listenAddr))

	http.ListenAndServe(s.listenAddr, router)
}

// func handlers
func (s *APIServer) handlerAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handlerGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handlerCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handlerDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// get account
func (s *APIServer) handlerGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// create account
func (s *APIServer) handlerCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// delete account
func (s *APIServer) handlerDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// transfer
func (s *APIServer) handlerTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
