package record

import (
	"strings"
	"testing"
	"time"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupTests() *gorm.DB {
	mocket.Catcher.Register()
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "user:test@tcp(127.0.0.1:3306)",
	}), &gorm.Config{})
	return db
}

func TestAdd(t *testing.T){
	db := SetupTests()
	tests := []struct {
		params []string
		expected	string
	}{
		{
			params: []string{},
			expected: addChoose,
		},
		{
			params: []string{"melon"},
			expected: strings.ReplaceAll(addSuccess, "<item>", "melon"),
		},
		{
			params: []string{"coconut", "pie"},
			expected: strings.ReplaceAll(addSuccess, "<item>", "coconut pie"),
		},
	}

	i := &Item{DB: db}
	for _, tt := range tests {
		actual, err := Add(i, tt.params)
		assert.Nil(t, err)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestShow(t *testing.T){
	db := SetupTests()
	i := &Item{DB: db}

	t.Run("no items", func(t *testing.T) {
		mocket.Catcher.Reset().NewMock().WithReply(nil)
		actual, err := Show(i, []string{"milk"})
		assert.Nil(t, err)
		assert.Equal(t, strings.ReplaceAll(itemNotExist, "<item>", "milk"), actual)
	})

	t.Run("no category", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "egg", "description": "Super tasty and cheap", "amount": 12, "unit": "piece(s)",
			"price": 98.50, "currency": "PHP", "expiration": time.Date(2021, 2, 26, 20, 34, 58, 651387237, time.UTC)}}
		mocket.Catcher.Reset().NewMock().WithReply(records)
		actual, err := Show(i, []string{"egg"})
		assert.Nil(t, err)
		assert.Equal(t, "*Egg* (_Uncategorized_)\n\nSuper tasty and cheap\nAmount: 12.00 piece(s)\nPrice: 98.50 PHP\nExpiration: 2021/02/26", actual)
	})

	t.Run("no description", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "egg", "amount": 12, "unit": "piece(s)", "category": "protein", "price": 98.50,
			"currency": "PHP", "expiration": time.Date(2021, 2, 26, 20, 34, 58, 651387237, time.UTC)}}
		mocket.Catcher.Reset().NewMock().WithReply(records)
		actual, err := Show(i, []string{"egg"})
		assert.Nil(t, err)
		assert.Equal(t, "*Egg* (Protein)\n\n_No description_\nAmount: 12.00 piece(s)\nPrice: 98.50 PHP\nExpiration: 2021/02/26", actual)
	})

	t.Run("no expiration", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "strawberry milk", "description": "Fruity", "amount": 2, "unit": "cup(s)",
			"category": "fruit", "price": 98.10, "currency": "PHP"}}
		mocket.Catcher.Reset().NewMock().WithReply(records)
		actual, err := Show(i, []string{"egg"})
		assert.Nil(t, err)
		assert.Equal(t, "*Strawberry Milk* (Fruit)\n\nFruity\nAmount: 2.00 cup(s)\nPrice: 98.10 PHP\nExpiration: _Not set_", actual)
	})
}

func TestList(t *testing.T) {
	db := SetupTests()
	i := &Item{DB: db}

	t.Run("invalid arguments", func(t *testing.T) {
		params := []string{"sort", "sort by", "filter", "filter by", "something made-up"}
		for _, p := range params {
			actual, err := List(i, []string{p})
			assert.Nil(t, err)
			assert.Equal(t, invalidListMsg, actual)
		}
	})
}

func TestUpdate(t *testing.T){
}

func TestDelete(t *testing.T){
}
