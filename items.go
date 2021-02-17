package main

import "gorm.io/gorm"

type Items struct {
	db *gorm.DB
}

func NewItems(db *gorm.DB) *Items {
	return &Items{db: db}
}
