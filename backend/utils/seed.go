package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"honya/backend/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedBooksAndReviews(db *gorm.DB) error {
	var bookCount int64
	if err := db.Model(&model.Book{}).Count(&bookCount).Error; err != nil {
		return err
	}

	if bookCount == 0 {
		var books []model.Book
		booksJSON, err := json.Marshal(BooksDummyData)
		if err != nil {
			return errors.New("failed to marshal books dummy data")
		}
		if err := json.Unmarshal(booksJSON, &books); err != nil {
			return errors.New("failed to parse hardcoded books data")
		}

		now := time.Now().Unix()
		for i := range books {
			if books[i].ID == uuid.Nil {
				books[i].ID = uuid.New()
			}
			books[i].CreatedAt = now
			books[i].UpdatedAt = now
		}

		if err := db.Create(&books).Error; err != nil {
			return err
		}

		fmt.Println("Seeded", len(books), "books successfully.")
	} else {
		fmt.Println("Books already exist in database. Skipping seeding.")
	}

	var reviewCount int64
	if err := db.Model(&model.Review{}).Count(&reviewCount).Error; err != nil {
		return err
	}

	if reviewCount > 0 {
		fmt.Println("Reviews already exist in database. Skipping seeding.")
		return nil
	}

	var reviews []model.Review
	// Convert ReviewsDummyData to JSON bytes first
	reviewsJSON, err := json.Marshal(ReviewsDummyData)
	if err != nil {
		return errors.New("failed to marshal reviews dummy data")
	}
	// Unmarshal the JSON bytes into reviews slice
	if err := json.Unmarshal(reviewsJSON, &reviews); err != nil {
		return errors.New("failed to parse hardcoded reviews data")
	}

	now := time.Now().Unix()
	for i := range reviews {
		if reviews[i].ID == uuid.Nil {
			reviews[i].ID = uuid.New()
		}
		if reviews[i].BookID == uuid.Nil {
			return fmt.Errorf("review at index %d does not have a valid BookID", i)
		}
		reviews[i].CreatedAt = now
		reviews[i].UpdatedAt = now
	}

	if err := db.Create(&reviews).Error; err != nil {
		return err
	}

	fmt.Println("Seeded", len(reviews), "reviews successfully.")
	return nil
}
