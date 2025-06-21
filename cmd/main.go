package main

import (
    "go-api-server/internal/routes"
    "log"
    "net/http"
)

func main() {
    router := routes.SetupRoutes()
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}