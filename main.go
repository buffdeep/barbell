package main

import (
	"log"
	"os"

	"bhavdeep.me/barbell/pkg/api"
)

func main() {
	servicePort := os.Getenv("PORT")
	if servicePort == "" {
		log.Fatal("Env variable $PORT not set. Aborting.")
	}
	api.HandleRequests(servicePort)
}
