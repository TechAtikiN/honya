package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title           string    `gorm:"type:varchar(255);not null;index:idx_book_title" json:"title"`
	Description     string    `gorm:"type:text" json:"description"`
	Category        string    `gorm:"type:varchar(100);index:idx_book_category" json:"category"`
	Image           string    `gorm:"type:varchar(255)" json:"image"`
	PublicationYear int       `gorm:"type:int;index:idx_book_publication_year" json:"publication_year"`
	Rating          float64   `gorm:"type:float;index:idx_book_rating" json:"rating"`
	Pages           int       `gorm:"type:int;index:idx_book_pages" json:"pages"`
	Isbn            string    `gorm:"type:varchar(20);unique;index:idx_book_isbn" json:"isbn"`
	CreatedAt       int64     `gorm:"autoCreateTime;index:idx_book_created_at" json:"created_at"`
	UpdatedAt       int64     `gorm:"autoUpdateTime;index:idx_book_updated_at" json:"updated_at"`
	AuthorName      string    `gorm:"type:varchar(100);index:idx_book_author_name" json:"author_name"`

	Reviews []Review `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"reviews,omitempty"`
}

func (Book) TableName() string {
	return "books"
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}

	return
}

func CreateBookIndexes(db *gorm.DB) error {
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_book_category_rating ON books (category, rating DESC)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_book_author_publication ON books (author_name, publication_year DESC)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_book_search ON books (title, author_name)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_book_fulltext ON books USING gin(to_tsvector('english', title || ' ' || description))").Error; err != nil {
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_book_rating_publication ON books (rating DESC, publication_year DESC)").Error; err != nil {
		return err
	}

	return nil
}
