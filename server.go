package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/wellingtonchida/searchstreet/internal/service"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type Server interface {
	Start() error
}

type APIServer struct {
	port string
	svc  service.Service
}

func NewServer(svc service.Service) *APIServer {
	port := os.Getenv("PORT")
	return &APIServer{
		port: port,
		svc:  svc,
	}
}

func (s *APIServer) Start() error {
	http.HandleFunc("/", makeHTTPHandler(s.FetchCep))
	log.Fatal(http.ListenAndServe(s.port, nil))

	return nil
}

func makeHTTPHandler(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "requestId", rand.Intn(100000))
		err := fn(ctx, w, r)
		if err != nil {
			log.Printf("Error handling request: %v", err)
			writeJson(w, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
}

func (s *APIServer) FetchCep(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	cep := r.URL.Query().Get("cep")
	//logic to get cep
	addressDb, err := s.svc.GetCepFromDb(ctx, cep)
	if err != nil {
		log.Printf("Error consulting CEP from db: %v", err)
		return err
	}
	if addressDb.Street != "" {
		return writeJson(w, http.StatusOK, addressDb)
	}

	address, err := s.svc.GetCep(ctx, cep)
	if err != nil {
		log.Printf("Error consulting CEP: %v", err)
		return err
	}

	address.ZipCode = strings.Replace(address.ZipCode, "-", "", -1)
	err = s.svc.InsertCep(ctx, address)
	if err != nil {
		log.Printf("Error inserting CEP: %v", err)
		return err
	}

	return writeJson(w, http.StatusOK, address)
}

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
