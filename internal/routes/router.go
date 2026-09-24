package routes

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-api-server/internal/handlers"
)

func SetupRoutes(pool *pgxpool.Pool) *mux.Router {
	posts := &handlers.PostHandler{DB: pool}

	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/posts", posts.List).Methods("GET")
	router.HandleFunc("/posts/{id}", posts.Get).Methods("GET")
	return router
}
