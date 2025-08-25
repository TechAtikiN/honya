package main

import (
	"fmt"
	"log"

	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/utils"
)

func main() {
	env, err := config.GetEnvConfig()
	if err != nil {
		log.Fatalf("Failed to get environment configuration: %v", err)
	}

	config.ConnectToDatabase(env.DatabaseURL)

	if err := utils.SeedBooksAndReviews(config.DB.Db); err != nil {
		log.Fatalf("Failed to seed books: %v", err)
	}

	fmt.Println("Database seeded successfully!")
}
