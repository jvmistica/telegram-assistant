package main

import (
	"strings"
	"testing"
	"time"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestShowItem(t *testing.T) {
	SetupTests()

	i := Items{db: DB}
	t.Run("no items", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "items" WHERE`).WithReply(nil)
		res, err := i.ShowItem([]string{"milk"})
		assert.Equal(t, "", res)
		assert.Equal(t, strings.ReplaceAll(itemNotExist, "<item>", "milk"), err.Error())
	})

	t.Run("one word", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "egg", "amount": 12, "unit": "piece(s)", "category": "protein", "price": 98.50, "currency": "PHP",
			"expiration": time.Date(2021, 2, 26, 20, 34, 58, 651387237, time.UTC)}}
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "items" WHERE`).WithReply(records)
		res, err := i.ShowItem([]string{"egg"})
		assert.Nil(t, err)
		assert.Equal(t, "*Egg* (Protein)\n\n_No description_\nAmount: 12.00 piece(s)\nPrice: 98.50 PHP\nExpiration: 2021/02/26", res)
	})

	t.Run("two words", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "strawberry milk", "amount": 2, "unit": "cup(s)", "category": "fruit", "price": 98.10, "currency": "PHP",
			"expiration": time.Date(2021, 2, 23, 20, 34, 58, 651387237, time.UTC)}}
		mocket.Catcher.Reset().NewMock().WithQuery(`SELECT * FROM "items" WHERE`).WithReply(records)
		res, err := i.ShowItem([]string{"egg"})
		assert.Nil(t, err)
		assert.Equal(t, "*Strawberry Milk* (Fruit)\n\n_No description_\nAmount: 2.00 cup(s)\nPrice: 98.10 PHP\nExpiration: 2021/02/23", res)
	})
}
