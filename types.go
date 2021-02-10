package main

import "time"

type Item struct {
	ID          uint
	Name        string
	Description string
	Amount      float32
	Unit        string
	Category    string
	Price       float32
	Currency    string
	Expiration  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
