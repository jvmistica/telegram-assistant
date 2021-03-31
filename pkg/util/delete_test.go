package util

import (
	"strings"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestDeleteItem(t *testing.T) {
	SetupTests()

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

	i := Items{db: DB}
	for _, tt := range tests {
		if tt.noRows {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			res, err := i.DeleteItem(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		} else {
			mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			res, err := i.DeleteItem(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		}
	}
}
