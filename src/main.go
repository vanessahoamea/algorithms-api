package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/vanessahoamea/algorithms-api/src/docs"
	"github.com/vanessahoamea/algorithms-api/src/handlers"
)

// @Title Algorithms API
// @Description A simple Go API that solves common Computer Science problems.
// @Version 1.0
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

	baseUrl := os.Getenv("BASE_URL")
	if port == "" {
		log.Fatal("Environment error: BASE_URL not found")
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

	// setting Swagger documentation route
	v1Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", baseUrl)),
	))

	// initializing server
	server := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Server starting on port %s...\n", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
