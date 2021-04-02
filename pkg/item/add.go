package item

import "strings"

// AddItem adds an item to the inventory
func (i *Items) AddItem(params []string) (string, error) {
	if len(params) == 0 {
		return addChoose, nil
	}

	item := strings.Join(params, " ")
	rec := Item{Name: item}
	err := i.db.Create(&rec)
	if err.Error != nil {
		return "", err.Error
	}

	msg := strings.ReplaceAll(addSuccess, "<item>", item)
	return msg, nil
}