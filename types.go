package main

import "time"

type Item struct {
	ID          uint
	Name        string
	Description string
	Amount      float32
	Category    string
	Price       float32
	Expiration  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
