package main

import (
	"errors"
	"strings"
)

// DeleteItem deletes an item
func (i *Items) DeleteItem(params []string) (string, error) {
	if len(params) == 0 {
		return deleteChoose, nil
	}

	var msg string
	item := strings.Join(params, " ")
	res := i.db.Where("name = ?", item).Delete(Item{})

	if res.RowsAffected == 0 {
		return "", errors.New(strings.ReplaceAll(itemNotExist, "<item>", item))
	}

	msg = strings.ReplaceAll(deleteSuccess, "<item>", item)

	return msg, nil
}
