package main

import (
	"strings"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupTests() *gorm.DB {
	mocket.Catcher.Register()
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "user:test@tcp(127.0.0.1:3306)",
	}), &gorm.Config{})
	DB = db
	return db
}

func TestAddItem(t *testing.T) {
	SetupTests()

	tests := []struct {
		params   []string
		expected string
	}{
		{
			params:   []string{},
			expected: addChoose,
		},
		{
			params:   []string{"melon"},
			expected: strings.ReplaceAll(addSuccess, "<item>", "melon"),
		},
		{
			params:   []string{"canned", "tuna"},
			expected: strings.ReplaceAll(addSuccess, "<item>", "canned tuna"),
		},
	}

	i := Items{db: DB}
	for _, tt := range tests {
		res, err := i.AddItem(tt.params)
		assert.Nil(t, err)
		assert.Equal(t, tt.expected, res)
	}
}
