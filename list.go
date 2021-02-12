package main

import (
	"fmt"
	"strings"
)

// ListItems returns all the items in the inventory
func (i *Items) ListItems(params string) string {
	var (
		items     []Item
		itemsList string
	)

	cmd := strings.Split(params, " ")
	switch cmd[0] {
	case "sort":
		i.db.Order(cmd[2]).Find(&items)
	case "filter":
		i.db.Where(fmt.Sprintf("%s %s '%s'", cmd[2], cmd[3], strings.Join(cmd[4:], " "))).Find(&items)
	default:
		i.db.Find(&items)
	}

	for _, item := range items {
		itemsList += item.Name + "\n"
	}

	return itemsList
}
