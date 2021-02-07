package main

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
		wantErr  bool
	}{
		{
			params:   []string{},
			expected: deleteChoose,
			wantErr:  false,
		},
		{
			params:   []string{"flour"},
			expected: strings.ReplaceAll(deleteSuccess, "<item>", "flour"),
			wantErr:  false,
		},
		{
			params:   []string{"almond", "flour"},
			expected: strings.ReplaceAll(deleteSuccess, "<item>", "almond flour"),
			wantErr:  false,
		},
		{
			params:   []string{"milk"},
			expected: strings.ReplaceAll(itemNotExist, "<item>", "milk"),
			wantErr:  true,
		},
	}

	i := Items{db: DB}
	for _, tt := range tests {
		if tt.wantErr {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			res, err := i.DeleteItem(tt.params)
			assert.Equal(t, "", res)
			assert.Equal(t, tt.expected, err.Error())
		} else {
			mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			res, err := i.DeleteItem(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		}
	}
}
