package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/vanessahoamea/algorithms-api/src/handlers"
)

func main() {
	// reading environment variables
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	if env == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Environment error: Failed to load environment variables")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Environment error: PORT not found")
	}

	// initializing router
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// setting v1 routes
	v1Router := chi.NewRouter()

	v1Router.Get("/status", handlers.HandleStatus)
	v1Router.Post("/n-queens", handlers.HandleNQueens)
	v1Router.Post("/knapsack", handlers.HandleKnapsack)
	v1Router.Post("/shortest-path", handlers.HandleShortestPath)

	router.Mount("/v1", v1Router)

	// initializing server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %s...\n", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
