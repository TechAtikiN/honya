package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/techatikin/backend/model"
	"gorm.io/gorm"
)

func SeedBooks(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.Book{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		// Books already exist, no seeding needed
		fmt.Println("Books already exist in database. Skipping seeding.")
		return nil
	}

	// Read seed JSON file
	file, err := os.ReadFile("scripts/seed/data/books.json")
	if err != nil {
		return errors.New("failed to read books.json")
	}

	var books []model.Book
	if err := json.Unmarshal(file, &books); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return errors.New("failed to parse books.json")
	}

	now := time.Now().Unix()
	for i := range books {
		// Generate a new UUID for each book instead of using the one from the JSON
		if books[i].ID == uuid.Nil {
			books[i].ID = uuid.New() // Generate new UUID
		}
		// Set CreatedAt and UpdatedAt to current Unix timestamp
		books[i].CreatedAt = now
		books[i].UpdatedAt = now
	}

	// Insert all books into the database
	if err := db.Create(&books).Error; err != nil {
		return err
	}

	fmt.Println("Seeded", len(books), "books successfully.")
	return nil
}
