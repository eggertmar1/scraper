package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"scraper/internal/db"
	"scraper/internal/models"       // Import models instead of repositories
	"scraper/internal/repositories" // Keep the repository import for interacting with the database

	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Initialize the database
	database, err := db.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer database.Close()
	// Initialize repository
	leagueRepo := repositories.NewLeagueRepository(database)

	// Example of inserting a new league (correctly using models.League)
	league := &models.League{
		Name:    "Premier League",
		Country: "England",
	}

	err = leagueRepo.Insert(context.Background(), league)
	if err != nil {
		log.Fatal("Failed to insert league:", err)
	} else {
		fmt.Printf("League inserted with ID: %d\n", league.ID)
	}

	// Start GraphQL server (you'll need gqlgen setup here)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Add GraphQL resolver setup here
	// Example: resolver := &graphql.Resolver{}

	// Serve GraphQL playground and API
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", handler.NewDefaultServer(resolver))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
