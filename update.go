package main

import "strings"

// UpdateItem updates an existing item's details
func (i *Items) UpdateItem(params []string) (string, error) {
	var msg string
	if len(params) == 0 {
		return updateChoose, nil
	}

	res := i.db.Model(&Item{}).Where("name = ?", params[1]).Update(params[0], strings.Join(params[2:], " "))
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected > 0 {
		msg = strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", params[1]), "<field>", params[0])
	} else {
		msg = strings.ReplaceAll(itemNotExist, "<item>", params[1])
	}

	return msg, nil
}
