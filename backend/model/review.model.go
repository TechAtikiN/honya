package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BookID    uuid.UUID `gorm:"type:uuid;not null;index:idx_review_book_id" json:"book_id"`
	Name      string    `gorm:"type:varchar(100);not null;index:idx_review_name" json:"name"`
	Email     string    `gorm:"type:varchar(100);not null;index:idx_review_email" json:"email"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt int64     `gorm:"autoCreateTime;index:idx_review_created_at" json:"created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime;index:idx_review_updated_at" json:"updated_at"`

	Book Book `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (Review) TableName() string {
	return "reviews"
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}

	return nil
}

func CreateReviewIndexes(db *gorm.DB) error {
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_review_book_created ON reviews (book_id, created_at DESC)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_review_email_created ON reviews (email, created_at DESC)").Error; err != nil {
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_review_created_book ON reviews (created_at DESC, book_id)").Error; err != nil {
		return err
	}

	return nil
}
