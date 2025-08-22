package main

import (
	"fmt"
	"log"

	"github.com/techatikin/backend/config"
	"github.com/techatikin/backend/utils"
)

func main() {
	config.ConnectToDatabase()

	if err := utils.SeedBooks(config.DB.Db); err != nil {
		log.Fatalf("Failed to seed books: %v", err)
	}

	fmt.Println("Database seeded successfully!")
}
