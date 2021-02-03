package main

import (
	"strings"
)

// CheckCommand checks if the command is valid and returns the appropriate response
func (i *Items) CheckCommand(cmd string) (string, error) {
	var msg string
	var params string

	if len(strings.Split(cmd, " ")) > 2 {
		params = strings.Join(strings.Split(cmd, " ")[2:], " ")
		cmd = strings.Join(strings.Split(cmd, " ")[:2], " ")
	}

	switch cmd {
	case "/start":
		msg = startMsg
	case "/list items":
		msg = i.ListItems()
		if msg == "" {
			msg = noItems
		}
	case "/add item":
		if params == "" {
			msg = addChoose
		} else {
			res, err := i.AddItem(params)
			if err != nil {
				return "", err
			}
			return res, nil
		}
	case "/edit item":
		msg = editChoose
	case "/delete item":
		if params == "" {
			msg = deleteChoose
		} else {
			res := i.DeleteItem(params)
			return res, nil
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
func (i *Items) ListItems() string {
	var (
		items     []Item
		itemsList string
	)

	i.db.Find(&items)
	for _, item := range items {
		itemsList += item.Name + "\n"
	}
	return itemsList
}

// AddItem adds an item to the inventory
func (i *Items) AddItem(item string) (string, error) {
	rec := Item{Name: item}

	err := i.db.Create(&rec)
	if err.Error != nil {
		return "", err.Error
	}

	msg := strings.ReplaceAll(addSuccess, "<item>", item)
	return msg, nil
}

// EditItem updates an item's properties
func (i *Items) EditItem(item, newItem string) {
	i.db.Model(&Item{}).Where("name = ?", item).Update("name", newItem)
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
