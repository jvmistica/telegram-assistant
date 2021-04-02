package item

import (
	"gorm.io/gorm"
	"time"
)

type Item struct {
	ID          uint
	Name        string
	Description string
	Amount      float32
	Unit        string
	Calories    uint16
	Category    string
	Price       float32
	Currency    string
	Expiration  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Items struct {
	db *gorm.DB
}

func NewItems(db *gorm.DB) *Items {
	return &Items{db: db}
}
