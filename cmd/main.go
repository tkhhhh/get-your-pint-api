package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go-api-server/internal/db"
	"go-api-server/internal/routes"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	pool, err := db.Connect(ctx, dbURL)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer pool.Close()

	if err := db.Migrate(ctx, pool); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := routes.SetupRoutes(pool)
	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
