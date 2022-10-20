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

const (
	fieldTag = "<field>"
)

func setupTestDB() *gorm.DB {
	mocket.Catcher.Register()
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DriverName: mocket.DriverName,
		DSN:        "user:test@tcp(127.0.0.1:3306)",
	}), &gorm.Config{})
	return db
}

func TestAdd(t *testing.T) {
	db := setupTestDB()
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "",
			expected: addChoose,
		},
		{
			name:     " ",
			expected: addChoose,
		},
		{
			name:     "melon",
			expected: strings.ReplaceAll(addSuccess, itemTag, "melon"),
		},
		{
			name:     "coconut pie",
			expected: strings.ReplaceAll(addSuccess, itemTag, "coconut pie"),
		},
		{
			name:     "thai chicken curry",
			expected: strings.ReplaceAll(addSuccess, itemTag, "thai chicken curry"),
		},
	}

	r := &RecordDB{DB: db}
	for _, tt := range tests {
		params := strings.Split(tt.name, " ")
		actual, err := r.Add(params)
		assert.Nil(t, err)
		assert.Equal(t, tt.expected, actual)
	}
}

func TestShow(t *testing.T) {
	db := setupTestDB()
	r := &RecordDB{DB: db}

	tests := map[string]struct {
		name     string
		record   []map[string]interface{}
		expected string
	}{
		"no record": {
			name:     "milk",
			record:   nil,
			expected: strings.ReplaceAll(itemNotExist, itemTag, "milk"),
		},
		"no category": {
			name: "egg",
			record: []map[string]interface{}{{"name": "egg", "description": "Super tasty and cheap", "amount": 12, "unit": "piece(s)",
				"price": 98.50, "currency": "PHP", "expiration": time.Date(2021, 2, 26, 20, 34, 58, 651387237, time.UTC)}},
			expected: "*egg* (_Uncategorized_)\n\nSuper tasty and cheap\nAmount: 12.00 piece(s)\nPrice: 98.50 PHP\nExpiration: 2021/02/26",
		},
		"no description": {
			name: "chocolate",
			record: []map[string]interface{}{{"name": "chocolate", "amount": 5, "unit": "bar(s)", "category": "snack", "price": 44.50,
				"currency": "PHP", "expiration": time.Date(2022, 3, 29, 20, 34, 58, 651387237, time.UTC)}},
			expected: "*chocolate* (snack)\n\n_No description_\nAmount: 5.00 bar(s)\nPrice: 44.50 PHP\nExpiration: 2022/03/29",
		},
		"no expiration": {
			name: "strawberry milk",
			record: []map[string]interface{}{{"name": "strawberry milk", "description": "Fruity", "amount": 2, "unit": "cup(s)",
				"category": "fruit", "price": 98.10, "currency": "PHP"}},
			expected: "*strawberry milk* (fruit)\n\nFruity\nAmount: 2.00 cup(s)\nPrice: 98.10 PHP\nExpiration: _Not set_",
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			mocket.Catcher.Reset().NewMock().WithReply(tt.record)
			actual, err := r.Show([]string{tt.name})
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestList(t *testing.T) {
	db := setupTestDB()
	r := &RecordDB{DB: db}

	t.Run("invalid arguments", func(t *testing.T) {
		params := []string{"sort", "sort by", "filter", "filter by", "something made-up"}
		for _, p := range params {
			actual, err := r.List([]string{p})
			assert.Nil(t, err)
			assert.Equal(t, invalidListMsg, actual)
		}
	})
}

func TestUpdate(t *testing.T) {
	db := setupTestDB()
	r := &RecordDB{DB: db}

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
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, itemTag, "melon"), fieldTag, "category"),
			wantErr:  false,
			noRows:   false,
		},
		{
			params:   []string{"melon", "amount", "2"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, itemTag, "melon"), fieldTag, "amount"),
			wantErr:  false,
			noRows:   false,
		},
		{
			params:   []string{"melon", "price", "30.50"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, itemTag, "melon"), fieldTag, "price"),
			wantErr:  false,
			noRows:   false,
		},
		{
			params:   []string{"egg", "amount", "12"},
			expected: strings.ReplaceAll(itemNotExist, itemTag, "egg"),
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
			actual, err := r.Update(tt.params)
			assert.Equal(t, "", actual)
			assert.Equal(t, tt.expected, err.Error())
		} else if tt.noRows {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			actual, err := r.Update(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		} else {
			mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			actual, err := r.Update(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		}
	}
}

func TestDelete(t *testing.T) {
	db := setupTestDB()
	r := &RecordDB{DB: db}

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
			expected: strings.ReplaceAll(deleteSuccess, itemTag, "flour"),
			noRows:   false,
		},
		{
			params:   []string{"almond", "flour"},
			expected: strings.ReplaceAll(deleteSuccess, itemTag, "almond flour"),
			noRows:   false,
		},
		{
			params:   []string{"milk"},
			expected: strings.ReplaceAll(itemNotExist, itemTag, "milk"),
			noRows:   true,
		},
	}

	for _, tt := range tests {
		if tt.noRows {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			actual, err := r.Delete(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		} else {
			mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			actual, err := r.Delete(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		}
	}
}
