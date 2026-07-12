package main

import (
	"log"
	"net/http"
	"os"

	"github.com/paulfwatts/my_stack/hello-world/internal/web"
)

func main() {
	server, err := web.NewServer()
	if err != nil {
		log.Fatalf("create server: %v", err)
	}

	address := ":8080"
	if value := os.Getenv("ADDR"); value != "" {
		address = value
	}

	log.Printf("listening on http://localhost%s", address)
	if err := http.ListenAndServe(address, server.Routes()); err != nil {
		log.Fatalf("listen and serve: %v", err)
	}
}
