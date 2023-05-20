package record

import (
	"strings"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	db := setupTestDB()
	r := &RecordDB{DB: db}

	t.Run("default response", func(t *testing.T) {
		record := []map[string]interface{}{{"name": "strawberry milk", "description": "Fruity", "amount": 2, "unit": "cup(s)",
			"category": "fruit", "price": 98.10, "currency": defaultCurrency}}
		mocket.Catcher.Reset().NewMock().WithReply(record)

		tests := []struct {
			data             string
			expectedResponse string
		}{
			{
				data:             "/start",
				expectedResponse: ResponseStart,
			},
			{
				data:             "/listitems",
				expectedResponse: "strawberry milk\n",
			},
			{
				data:             "/showitem",
				expectedResponse: ResponseShow,
			},
			{
				data:             "/additem",
				expectedResponse: ResponseAdd,
			},
			{
				data:             "/updateitem",
				expectedResponse: ResponseUpdate,
			},
			{
				data:             "/deleteitem",
				expectedResponse: ResponseDelete,
			},
			{
				data:             "random string",
				expectedResponse: ResponseInvalid,
			},
		}

		for _, tt := range tests {
			actual, err := r.CheckCommand(tt.data)
			assert.Equal(t, tt.expectedResponse, actual)
			assert.Nil(t, err)
		}
	})

	t.Run("with params response", func(t *testing.T) {
		record := []map[string]interface{}{
			{"name": "strawberry milk", "description": "Fruity", "amount": 2, "unit": "cup(s)",
				"category": "fruit", "price": 98.10, "currency": defaultCurrency},
			{"name": "chocolate", "amount": 5, "unit": "bar(s)", "category": "snack", "price": 44.50, "currency": defaultCurrency}}
		mocket.Catcher.Reset().NewMock().WithReply(record)

		tests := []struct {
			data             string
			expectedResponse string
			updateOrDelete   bool
		}{
			{
				data:             "/showitem strawberry milk",
				expectedResponse: "*strawberry milk* (fruit)\n\nFruity\nAmount: 2.00 cup(s)\nPrice: 98.10 EUR\nExpiration: _Not set_",
			},
			{
				data:             "/additem chocolate",
				expectedResponse: strings.ReplaceAll(ResponseSuccessAdd, itemTag, "chocolate"),
			},
			{
				data:             "/updateitem chocolate description sweet",
				expectedResponse: strings.ReplaceAll(strings.ReplaceAll(ResponseSuccessUpdate, itemTag, "chocolate"), fieldTag, "description"),
				updateOrDelete:   true,
			},
			{
				data:             "/deleteitem chocolate",
				expectedResponse: strings.ReplaceAll(ResponseSuccessDelete, itemTag, "chocolate"),
				updateOrDelete:   true,
			},
		}

		for _, tt := range tests {
			if tt.updateOrDelete {
				mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			}
			actual, err := r.CheckCommand(tt.data)
			assert.Equal(t, tt.expectedResponse, actual)
			assert.Nil(t, err)
		}
	})
}
