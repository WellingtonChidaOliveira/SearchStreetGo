package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wellingtonchida/searchstreet/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	serviceCep := service.Init()
	svc := NewLogginService(serviceCep)
	api := NewServer(svc)
	api.Start()

}
