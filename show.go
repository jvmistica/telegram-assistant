package main

import (
	"errors"
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
		return "", errors.New(strings.ReplaceAll(itemNotExist, "<item>", strings.Join(param, " ")))
	}

	details = fmt.Sprintf("*%s* (%s)\n%s\nAmount: %f %s\nPrice: %f %s\nExpiration: %s",
		strings.Title(item.Name), strings.Title(item.Category), item.Description, item.Amount, item.Unit, item.Price, item.Currency, item.Expiration)

	return details, nil
}
