package main

import "strings"

// DeleteItem deletes an item
func (i *Items) DeleteItem(params []string) string {
	if len(params) == 0 {
		return deleteChoose
	}

	var msg string
	item := strings.Join(params, " ")
	res := i.db.Where("name = ?", item).Delete(Item{})
	if res.RowsAffected > 0 {
		msg = strings.ReplaceAll(deleteSuccess, "<item>", item)
	} else {
		msg = strings.ReplaceAll(itemNotExist, "<item>", item)
	}

	return msg
}
