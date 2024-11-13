package main

import (
	"fmt"
	"log"
	"mctest-agent/internal/http"
)

func main() {
    // Log the start of the application
    log.SetFlags(0) // Remove detalhes de timestamp, etc.
    log.Println("Starting mctest-agent...")

    err := http.Init()

	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error initializing HTTP server: %v", err)
    } else {
        log.Println("HTTP server started successfully")
    }
    // Adiciona a linha separadora no final do arquivo
    log.Println("\n=================\n")
}
