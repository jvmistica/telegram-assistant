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

func TestAdd(t *testing.T) {
	db := SetupTests()
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
			params:   []string{"coconut", "pie"},
			expected: strings.ReplaceAll(addSuccess, "<item>", "coconut pie"),
		},
	}

	i := &RecordDB{DB: db}
	for _, tt := range tests {
		actual, err := Add(i, tt.params)
		assert.Nil(t, err)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestShow(t *testing.T) {
	db := SetupTests()
	i := &RecordDB{DB: db}

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
		assert.Equal(t, "*egg* (_Uncategorized_)\n\nSuper tasty and cheap\nAmount: 12.00 piece(s)\nPrice: 98.50 PHP\nExpiration: 2021/02/26", actual)
	})

	t.Run("no description", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "egg", "amount": 12, "unit": "piece(s)", "category": "protein", "price": 98.50,
			"currency": "PHP", "expiration": time.Date(2021, 2, 26, 20, 34, 58, 651387237, time.UTC)}}
		mocket.Catcher.Reset().NewMock().WithReply(records)
		actual, err := Show(i, []string{"egg"})
		assert.Nil(t, err)
		assert.Equal(t, "*egg* (protein)\n\n_No description_\nAmount: 12.00 piece(s)\nPrice: 98.50 PHP\nExpiration: 2021/02/26", actual)
	})

	t.Run("no expiration", func(t *testing.T) {
		records := []map[string]interface{}{{"name": "strawberry milk", "description": "Fruity", "amount": 2, "unit": "cup(s)",
			"category": "fruit", "price": 98.10, "currency": "PHP"}}
		mocket.Catcher.Reset().NewMock().WithReply(records)
		actual, err := Show(i, []string{"egg"})
		assert.Nil(t, err)
		assert.Equal(t, "*strawberry milk* (fruit)\n\nFruity\nAmount: 2.00 cup(s)\nPrice: 98.10 PHP\nExpiration: _Not set_", actual)
	})
}

func TestList(t *testing.T) {
	db := SetupTests()
	i := &RecordDB{DB: db}

	t.Run("invalid arguments", func(t *testing.T) {
		params := []string{"sort", "sort by", "filter", "filter by", "something made-up"}
		for _, p := range params {
			actual, err := List(i, []string{p})
			assert.Nil(t, err)
			assert.Equal(t, invalidListMsg, actual)
		}
	})
}

func TestUpdate(t *testing.T) {
	db := SetupTests()
	i := &RecordDB{DB: db}

	tests := []struct {
		params   []string
		expected string
		wantErr  bool
		noRows   bool
	}{
		{
			params:   []string{},
			expected: updateChoose,
			wantErr:  false,
			noRows:   false,
		},
		{
			params:   []string{"melon", "category", "fruit"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", "melon"), "<field>", "category"),
			wantErr:  false,
			noRows:   false,
		},
		{
			params:   []string{"melon", "amount", "2"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", "melon"), "<field>", "amount"),
			wantErr:  false,
			noRows:   false,
		},
		{
			params:   []string{"melon", "price", "30.50"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", "melon"), "<field>", "price"),
			wantErr:  false,
			noRows:   false,
		},
		{
			params:   []string{"egg", "amount", "12"},
			expected: strings.ReplaceAll(itemNotExist, "<item>", "egg"),
			wantErr:  false,
			noRows:   true,
		},
		{
			params:   []string{"melon"},
			expected: updateInvalid,
			wantErr:  true,
			noRows:   false,
		},
		{
			params:   []string{"melon", "price"},
			expected: updateInvalid,
			wantErr:  true,
			noRows:   false,
		},
	}

	for _, tt := range tests {
		if tt.wantErr {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			actual, err := Update(i, tt.params)
			assert.Equal(t, "", actual)
			assert.Equal(t, tt.expected, err.Error())
		} else if tt.noRows {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			actual, err := Update(i, tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		} else {
			mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			actual, err := Update(i, tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		}
	}
}

func TestDelete(t *testing.T) {
	db := SetupTests()
	i := &RecordDB{DB: db}

	tests := []struct {
		params   []string
		expected string
		noRows   bool
	}{
		{
			params:   []string{},
			expected: deleteChoose,
			noRows:   false,
		},
		{
			params:   []string{"flour"},
			expected: strings.ReplaceAll(deleteSuccess, "<item>", "flour"),
			noRows:   false,
		},
		{
			params:   []string{"almond", "flour"},
			expected: strings.ReplaceAll(deleteSuccess, "<item>", "almond flour"),
			noRows:   false,
		},
		{
			params:   []string{"milk"},
			expected: strings.ReplaceAll(itemNotExist, "<item>", "milk"),
			noRows:   true,
		},
	}

	for _, tt := range tests {
		if tt.noRows {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			actual, err := Delete(i, tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		} else {
			mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			actual, err := Delete(i, tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		}
	}
}
