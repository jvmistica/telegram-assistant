package main

import (
	"fmt"
	"strings"
)

// AddItemDetails updates an existing item's details
func (i *Items) AddItemDetails(params []string) (string, error) {
	var err error

	if len(params) == 0 {
		return addChoose, nil
	}

	switch params[0] {
	case "description":
		err = i.EditItem("description", params[1], strings.Join(params[2:], " "))
	case "amount":
		err = i.EditItem("amount", params[1], strings.Join(params[2:], " "))
	case "category":
		err = i.EditItem("category", params[1], strings.Join(params[2:], " "))
	case "price":
		err = i.EditItem("price", params[1], strings.Join(params[2:], " "))
	case "expiration":
		err = i.EditItem("expiration", params[1], strings.Join(params[2:], " "))
	default:
		res, err := i.AddItem(params[0])
		if err != nil {
			return "", err
		}
		return res, nil
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Updated %s's %s.", params[1], params[0]), nil
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
