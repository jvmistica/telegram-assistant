package record

import (
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	db := setupTestDB()
	r := &RecordDB{DB: db}
	record := []map[string]interface{}{{"name": "strawberry milk", "description": "Fruity", "amount": 2, "unit": "cup(s)",
		"category": "fruit", "price": 98.10, "currency": "PHP"}}
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
	}

	for _, tt := range tests {
		actual, err := r.CheckCommand(tt.data)
		assert.Equal(t, tt.expectedResponse, actual)
		assert.Nil(t, err)
	}
}
