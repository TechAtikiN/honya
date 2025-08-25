package main

import (
	"fmt"
	"honya/backend/config"
	"honya/backend/utils"
	"log"
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
