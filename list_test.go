package main

import (
	"testing"
	"time"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestListItems(t *testing.T) {
	SetupTests()

	i := Items{db: DB}
	t.Run("no items", func(t *testing.T) {
		res, err := i.ListItems("")
		assert.Nil(t, err)
		assert.Equal(t, "There are no items in your inventory.", res)
	})

	t.Run("has items", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "egg", "amount": 12, "category": "protein", "price": 98.50, "expiration": time.Date(2021, 2, 26, 20, 34, 58, 651387237, time.UTC)},
			{"name": "oil", "amount": 1, "category": "fat", "price": 20.00, "expiration": time.Date(2021, 5, 23, 20, 34, 58, 651387237, time.UTC)},
			{"name": "flour", "amount": 2, "category": "carbohydrate", "price": 18.68, "expiration": time.Date(2021, 8, 22, 20, 34, 58, 651387237, time.UTC)}}
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "items"`).WithReply(records)
		res, err := i.ListItems("")
		assert.Nil(t, err)
		assert.Equal(t, "egg\noil\nflour\n", res)
	})
}
