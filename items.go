package main

import (
	"fmt"
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
		msg = "Welcome. Here's a list of my commands:\n" +
			"/list items	- List items in your inventory\n" +
			"/add item	- Add an item to your inventory\n" +
			"/edit item	- Edit an item in your inventory\n" +
			"/delete item	- Delete an item in your inventory\n\n" +
			"To view the full list of commands and sub-commands, enter /list commands"
	case "/list items":
		msg = i.ListItems()
		if msg == "" {
			msg = "There are no items in your inventory."
		}
	case "/add item":
		if params == "" {
			msg = "What is this item called?"
		} else {
			res, err := i.AddItem(params)
			if err != nil {
				return "", err
			}
			return res, nil
		}
	case "/edit item":
		msg = "Which item do you want to edit?"
	case "/delete item":
		if params == "" {
			msg = "Which item do you want to delete?"
		} else {
			res := i.DeleteItem(params)
			return res, nil
		}
	default:
		msg = "That's not a valid command. Here's a list of valid commands:\n" +
			"/list items	- List items in your inventory\n" +
			"/add item	- Add an item to your inventory\n" +
			"/edit item	- Edit an item in your inventory\n" +
			"/delete item	- Delete an item in your inventory\n"
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

	msg := fmt.Sprintf("Added \"%s\" to the inventory.\n\n", item) +
		"Add more details about this item using the commands below:\n" +
		fmt.Sprintf("/add item description %s <description>\n", item) +
		fmt.Sprintf("/add item amount %s <amount>\n", item) +
		fmt.Sprintf("/add item category %s <category>\n", item) +
		fmt.Sprintf("/add item price %s <price>\n", item) +
		fmt.Sprintf("/add item expiration %s <expiration>\n", item)
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
		msg = fmt.Sprintf("Removed \"%s\" from the inventory.", item)
	} else {
		msg = fmt.Sprintf("Item \"%s\" does not exist in the inventory.", item)
	}

	return msg
}
