package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Title           string    `gorm:"type:varchar(255);not null" json:"title"`
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Description     string    `gorm:"type:text" json:"description"`
	Category        string    `gorm:"type:varchar(100)" json:"category"`
	Image           string    `gorm:"type:varchar(255)" json:"image"`
	PublicationYear int       `gorm:"type:int" json:"publication_year"`
	Rating          float64   `gorm:"type:float" json:"rating"`
	Pages           int       `gorm:"type:int" json:"pages"`
	Isbn            string    `gorm:"type:varchar(20);unique" json:"isbn"`
	CreatedAt       int64     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       int64     `gorm:"autoUpdateTime" json:"updated_at"`
	AuthorName      string    `gorm:"type:varchar(100)" json:"author_name"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return
}
