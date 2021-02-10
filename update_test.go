package main

import (
	"strings"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func TestUpdateItem(t *testing.T) {
	SetupTests()

	tests := []struct {
		params   []string
		expected string
		wantErr  bool
	}{
		{
			params:   []string{},
			expected: updateChoose,
			wantErr:  false,
		},
		{
			params:   []string{"melon", "category", "fruit"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", "melon"), "<field>", "category"),
			wantErr:  false,
		},
		{
			params:   []string{"melon", "amount", "2"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", "melon"), "<field>", "amount"),
			wantErr:  false,
		},
		{
			params:   []string{"melon", "price", "30.50"},
			expected: strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", "melon"), "<field>", "price"),
			wantErr:  false,
		},
		{
			params:   []string{"melon"},
			expected: updateInvalid,
			wantErr:  true,
		},
		{
			params:   []string{"melon", "price"},
			expected: updateInvalid,
			wantErr:  true,
		},
		{
			params:   []string{"egg", "amount", "12"},
			expected: strings.ReplaceAll(itemNotExist, "<item>", "egg"),
			wantErr:  true,
		},
	}

	i := Items{db: DB}
	for _, tt := range tests {
		if tt.wantErr {
			mocket.Catcher.Reset().NewMock().WithRowsNum(0)
			res, err := i.UpdateItem(tt.params)
			assert.Equal(t, "", res)
			assert.Equal(t, tt.expected, err.Error())
		} else {
			mocket.Catcher.Reset().NewMock().WithRowsNum(1)
			res, err := i.UpdateItem(tt.params)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		}
	}
}
