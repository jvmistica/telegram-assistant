package recipe

import (
	"gorm.io/gorm"
	"time"
)

type Recipe struct {
	ID           uint
	Name         string
	Description  string
	Ingredients  string
	Instructions string
	Category     string
	Servings     uint
	Price        float32
	Currency     string
	Expiration   time.Time
	Calories     uint16
}

type Recipes struct {
	db *gorm.DB
}

func NewRecipes(db *gorm.DB) *Recipes {
	return &Recipes{db: db}
}
