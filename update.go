package main

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// UpdateItem updates an existing item's details
func (i *Items) UpdateItem(params []string) (string, error) {
	var (
		msg string
		res *gorm.DB
	)

	if len(params) == 0 {
		return updateChoose, nil
	}

	if len(params) < 3 {
		return "", errors.New(updateInvalid)
	}

	if params[1] == "amount" && len(params) > 3 {
		f, err := strconv.ParseFloat(params[2], 32)
		if err != nil {
			return "", err
		}

		res = i.db.Model(&Item{}).Where("name = ?", params[0]).Updates(Item{Amount: float32(f), Unit: params[3]})
		if res.Error != nil {
			return "", res.Error
		}
	} else if params[1] == "price" && len(params) > 3 {
		f, err := strconv.ParseFloat(params[2], 32)
		if err != nil {
			return "", err
		}

		res = i.db.Model(&Item{}).Where("name = ?", params[0]).Updates(Item{Price: float32(f), Currency: params[3]})
		if res.Error != nil {
			return "", res.Error
		}
	} else {
		res = i.db.Model(&Item{}).Where("name = ?", params[0]).Update(params[1], strings.Join(params[2:], " "))
		if res.Error != nil {
			return "", res.Error
		}
	}

	if res.RowsAffected == 0 {
		msg = strings.ReplaceAll(itemNotExist, "<item>", params[0])
	} else {
		msg = strings.ReplaceAll(strings.ReplaceAll(updateSuccess, "<item>", params[0]), "<field>", params[1])
	}

	return msg, nil
}
