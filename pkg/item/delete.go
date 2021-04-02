package item

import "strings"

// DeleteItem deletes an item
func (i *Items) DeleteItem(params []string) (string, error) {
	var msg string

	if len(params) == 0 {
		return deleteChoose, nil
	}

	item := strings.Join(params, " ")
	res := i.db.Where("name = ?", item).Delete(Item{})
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected == 0 {
		msg = strings.ReplaceAll(itemNotExist, "<item>", item)
	} else {
		msg = strings.ReplaceAll(deleteSuccess, "<item>", item)
	}

	return msg, nil
}
