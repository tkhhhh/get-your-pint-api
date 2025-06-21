package routes

import (
    "github.com/gorilla/mux"
    "go-api-server/internal/handlers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
    router.HandleFunc("/posts", handlers.GetPosts).Methods("GET") // Endpoint for blog posts
    router.HandleFunc("/posts/{id}", handlers.GetPostByID).Methods("GET") // Endpoint for a single post
    return router
}