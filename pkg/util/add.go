package util

import (
	"strings"
	"github.com/jvmistica/henchmaid/pkg/types"
)

// AddItem adds an item to the inventory
func (i *Items) AddItem(params []string) (string, error) {
	if len(params) == 0 {
		return addChoose, nil
	}

	item := strings.Join(params, " ")
	rec := types.Item{Name: item}
	err := i.db.Create(&rec)
	if err.Error != nil {
		return "", err.Error
	}

	msg := strings.ReplaceAll(addSuccess, "<item>", item)
	return msg, nil
}
