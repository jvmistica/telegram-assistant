package main

import (
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

func TestCheckCommand(t *testing.T) {
	SetupTests()

	tests := []struct {
		cmd      string
		expected string
	}{
		{
			"/start",
			"Welcome",
		},
		{
			"/list items",
			"no items in your inventory",
		},
		{
			"/add item",
			"What is this item called?",
		},
		{
			"/edit item",
			"Which item do you want to edit?",
		},
		{
			"/delete item",
			"Which item do you want to delete?",
		},
		{
			"/end",
			"not a valid command",
		},
		{
			"gibberish",
			"not a valid command",
		},
	}

	i := Items{db: DB}
	for _, test := range tests {
		result := i.CheckCommand(test.cmd)
		assert.Contains(t, result, test.expected)
	}
}

func TestCheckItem(t *testing.T) {
	SetupTests()

	tests := []struct {
		itemName string
		sqlResp  []map[string]interface{}
		expected bool
	}{
		{
			itemName: "onion",
			sqlResp:  []map[string]interface{}{},
			expected: false,
		},
		{
			itemName: "milk",
			sqlResp:  []map[string]interface{}{{"name": "milk"}},
			expected: true,
		},
		{
			itemName: "strawberry",
			sqlResp:  []map[string]interface{}{{"id": 123, "name": "strawberry"}, {"id": 234, "name": "strawberry"}, {"id": 456, "name": "strawberry"}},
			expected: true,
		},
	}

	i := Items{db: DB}
	for _, test := range tests {
		mocket.Catcher.Reset().NewMock().WithReply(test.sqlResp)
		res := i.CheckItem(test.itemName)
		assert.Equal(t, test.expected, res)
	}
}

func TestListItems(t *testing.T) {
	SetupTests()

	tests := []struct {
		sqlResp  []map[string]interface{}
		expected string
	}{
		{
			sqlResp:  []map[string]interface{}{},
			expected: "",
		},
		{
			sqlResp:  []map[string]interface{}{{"name": "TEST"}},
			expected: "TEST\n",
		},
		{
			sqlResp:  []map[string]interface{}{{"name": "TEST"}, {"name": "ANOTHER"}},
			expected: "TEST\nANOTHER\n",
		},
	}

	i := Items{db: DB}
	for _, test := range tests {
		mocket.Catcher.Reset().NewMock().WithReply(test.sqlResp)
		res := i.ListItems()
		assert.Equal(t, test.expected, res)
	}
}

// TODO: Improve this function/test
func TestAddItem(t *testing.T) {
	SetupTests()

	i := Items{db: DB}
	mocket.Catcher.Reset().NewMock().WithReply(nil)
	res := i.AddItem("chocolate")
	assert.Nil(t, res)
}

// TODO: Improve this function/test
func TestEditItem(t *testing.T) {
	SetupTests()

	i := Items{db: DB}
	i.EditItem("chocolate", "milk chocolate")
}

// TODO: Improve this function/test
func TestDeleteItem(t *testing.T) {
	SetupTests()

	i := Items{db: DB}
	sqlResp := []map[string]interface{}{{"name": "milk chocolate"}}
	mocket.Catcher.Reset().NewMock().WithReply(sqlResp)
	i.DeleteItem("milk chocolate")
}
