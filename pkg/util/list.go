package util

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"github.com/jvmistica/henchmaid/pkg/types"
)

// ListItems returns all the items in the inventory
func (i *Items) ListItems(params string) (string, error) {
	var (
		items     []types.Item
		itemsList string
		res       *gorm.DB
	)

	cmd := strings.Split(params, " ")
	if cmd[0] == "" && len(cmd) == 1 {
		res = i.db.Find(&items)
	} else if len(cmd) >= 3 && strings.Join(cmd[:2], " ") == "sort by" {
		if len(cmd) > 3 && (cmd[3] == "asc" || cmd[3] == "desc") {
			res = i.db.Order(fmt.Sprintf("%s %s", cmd[2], cmd[3])).Find(&items)
		} else {
			res = i.db.Order(cmd[2]).Find(&items)
		}
	} else if len(cmd) >= 5 && strings.Join(cmd[:2], " ") == "filter by" {
		res = i.db.Where(fmt.Sprintf("%s %s '%s'", cmd[2], cmd[3], strings.Join(cmd[4:], " "))).Find(&items)
		if res.RowsAffected == 0 {
			itemsList = noMatchFilter
		}
	} else {
		itemsList = invalidListMsg
	}

	if res != nil && res.Error != nil {
		return "", res.Error
	}

	for _, item := range items {
		if len(cmd) >= 3 && strings.Join(cmd[:2], " ") == "sort by" {
			if item.Expiration.Year() == 0 {
				itemsList += fmt.Sprintf("Not Available - %s\n", item.Name)
			} else {
				itemsList += fmt.Sprintf("%s %d - %s\n", item.Expiration.Month().String()[:3], item.Expiration.Year(), item.Name)
			}
		} else {
			itemsList += item.Name + "\n"
		}
	}

	if itemsList == "" {
		itemsList = noItems
	}

	return itemsList, nil
}
