package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BookID    uuid.UUID `gorm:"type:uuid;not null" json:"book_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);not null" json:"email"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime" json:"updated_at"`

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
