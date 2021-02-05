package main

import "strings"

// CheckCommand checks if the command is valid and returns the appropriate response
func (i *Items) CheckCommand(cmd string) (string, error) {
	var msg string
	var params []string

	if len(strings.Split(cmd, " ")) > 2 {
		params = strings.Split(cmd, " ")[2:]
		cmd = strings.Join(strings.Split(cmd, " ")[:2], " ")
	}

	switch cmd {
	case "/start":
		msg = startMsg
	case "/list items":
		if len(params) == 0 {
			msg = i.ListItems("")
			if msg == "" {
				msg = noItems
			}
		} else {
			msg = i.ListItems(strings.Join(params[0:], " "))
		}
	case "/add item":
		msg, err := i.AddItemDetails(params)
		if err != nil {
			return "", err
		}
		return msg, nil
	case "/edit item":
		msg = editChoose
	case "/delete item":
		if len(params) == 0 {
			msg = deleteChoose
		} else {
			msg = i.DeleteItem(params[0])
		}
	default:
		msg = invalidMsg
	}
	return msg, nil
}

// CheckItem returns true if item exists
func (i *Items) CheckItem(item string) bool {
	var rec Item
	res := i.db.Where("name = ?", item).Find(&rec)
	return res.RowsAffected == 1
}

// ListItems returns all the items in the inventory
func (i *Items) ListItems(params string) string {
	var (
		items     []Item
		itemsList string
	)

	if params == "sort by name" {
		i.db.Order("name").Find(&items)
	} else {
		i.db.Find(&items)
	}

	for _, item := range items {
		itemsList += item.Name + "\n"
	}

	return itemsList
}

// EditItem updates an item's properties
func (i *Items) EditItem(field, item, newItem string) error {
	res := i.db.Model(&Item{}).Where("name = ?", item).Update(field, newItem)
	return res.Error
}

// DeleteItem deletes an item
func (i *Items) DeleteItem(item string) string {
	var msg string
	res := i.db.Where("name = ?", item).Delete(Item{})
	if res.RowsAffected > 0 {
		msg = strings.ReplaceAll(deleteSuccess, "<item>", item)
	} else {
		msg = strings.ReplaceAll(itemNotExist, "<item>", item)
	}

	return msg
}
