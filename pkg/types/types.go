package types

import "time"

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
