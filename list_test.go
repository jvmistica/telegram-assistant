package main

import (
	"testing"
	"time"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestListItemsInvalidArgs(t *testing.T) {
	params := []string{"sort", "sort by", "filter", "filter by", "something made-up"}

	i := Items{}
	for _, p := range params {
		res, err := i.ListItems(p)
		assert.Equal(t, invalidListMsg, res)
		assert.Nil(t, err)
	}
}

func TestListItems(t *testing.T) {
	SetupTests()

	i := Items{db: DB}
	t.Run("no items", func(t *testing.T) {
		res, err := i.ListItems("")
		assert.Nil(t, err)
		assert.Equal(t, noItems, res)
	})

	t.Run("sort no items", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithReply(nil)
		res, err := i.ListItems("sort by name")
		assert.Nil(t, err)
		assert.Equal(t, noItems, res)
	})

	t.Run("filter no match", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithReply(nil)
		res, err := i.ListItems("filter by amount = 999.99")
		assert.Nil(t, err)
		assert.Equal(t, noMatchFilter, res)
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
