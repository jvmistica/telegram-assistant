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
			expectedResponse: startMsg,
		},
		{
			data:             "/listitems",
			expectedResponse: "strawberry milk\n",
		},
		{
			data:             "/showitem",
			expectedResponse: showChoose,
		},
		{
			data:             "/additem",
			expectedResponse: addChoose,
		},
		{
			data:             "/updateitem",
			expectedResponse: updateChoose,
		},
		{
			data:             "/deleteitem",
			expectedResponse: deleteChoose,
		},
	}

	for _, tt := range tests {
		actual, err := r.CheckCommand(tt.data)
		assert.Equal(t, tt.expectedResponse, actual)
		assert.Nil(t, err)
	}
}
