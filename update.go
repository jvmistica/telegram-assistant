package main

import (
	"errors"
	"strings"
)

// UpdateItem updates an existing item's details
func (i *Items) UpdateItem(params []string) (string, error) {
	var msg string
	if len(params) == 0 {
		return updateChoose, nil
	}

	if len(params) < 3 {
		return "", errors.New(updateInvalid)
	}

	res := i.db.Model(&Item{}).Where("name = ?", params[0]).Update(params[1], strings.Join(params[2:], " "))
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected == 0 {
		return "", errors.New(strings.ReplaceAll(itemNotExist, "<item>", params[0]))
	}

	msg = strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", params[0]), "<field>", params[1])

	return msg, nil
}
