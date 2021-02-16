package main

import (
	"fmt"
	"strings"
)

// ShowItem returns the details of an item
func (i *Items) ShowItem(param []string) (string, error) {
	var (
		item    Item
		details string
	)

	res := i.db.Where("name = ?", strings.Join(param, " ")).Find(&item)
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected == 0 {
		return strings.ReplaceAll(itemNotExist, "<item>", strings.Join(param, " ")), nil
	}

	category := "_Uncategorized_"
	if item.Category != "" {
		category = strings.Title(item.Category)
	}

	description := "_No description_"
	if item.Description != "" {
		description = item.Description
	}

	exp := item.Expiration
	expiration := exp.Format("2006/01/02")
	if expiration[0] == '0' {
		expiration = "_Not set_"
	}

	details = fmt.Sprintf("*%s* (%s)\n\n%s\nAmount: %.2f %s\nPrice: %.2f %s\nExpiration: %s",
		strings.Title(item.Name), category, description, item.Amount, item.Unit, item.Price, item.Currency, expiration)

	return details, nil
}
