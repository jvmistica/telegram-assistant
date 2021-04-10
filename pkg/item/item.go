package item

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ItemRecord struct {
	ID          uint
	Name        string
	Description string
	Amount      float32
	Unit        string
	Calories    uint16
	Category    string
	Price       float32
	Currency    string
	Expiration  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Item struct {
	DB *gorm.DB
}

func (i *Item) AddRecord(record string) error {
	rec := ItemRecord{Name: record}
	err := i.DB.Create(&rec)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (i *Item) ShowRecord(record string) (string, error) {
	var (
		item    ItemRecord
		details string
	)

	res := i.DB.Where("name = ?", record).Find(&item)
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected == 0 {
		return strings.ReplaceAll(itemNotExist, "<item>", record), nil
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

func (i *Item) ListRecords(cmd []string) (string, error) {
	var (
		items     []ItemRecord
		itemsList string
		res       *gorm.DB
	)

	if len(cmd) <= 1 {
		res = i.DB.Find(&items)
	} else if len(cmd) >= 3 && strings.Join(cmd[:2], " ") == "sort by" {
		if len(cmd) > 3 && (cmd[3] == "asc" || cmd[3] == "desc") {
			res = i.DB.Order(fmt.Sprintf("%s %s", cmd[2], cmd[3])).Find(&items)
		} else {
			res = i.DB.Order(cmd[2]).Find(&items)
		}
	} else if len(cmd) >= 5 && strings.Join(cmd[:2], " ") == "filter by" {
		res = i.DB.Where(fmt.Sprintf("%s %s '%s'", cmd[2], cmd[3], strings.Join(cmd[4:], " "))).Find(&items)
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
		itemsList = noDB
	}

	return itemsList, nil
}

func (i *Item) UpdateRecord(params []string) (string, error) {
	var (
		msg string
		res *gorm.DB
	)

	if params[1] == "amount" && len(params) > 3 {
		f, err := strconv.ParseFloat(params[2], 32)
		if err != nil {
			return "", err
		}

		res = i.DB.Model(&ItemRecord{}).Where("name = ?", params[0]).Updates(ItemRecord{Amount: float32(f), Unit: params[3]})
		if res.Error != nil {
			return "", res.Error
		}
	} else if params[1] == "price" && len(params) > 3 {
		f, err := strconv.ParseFloat(params[2], 32)
		if err != nil {
			return "", err
		}

		res = i.DB.Model(&ItemRecord{}).Where("name = ?", params[0]).Updates(ItemRecord{Price: float32(f), Currency: params[3]})
		if res.Error != nil {
			return "", res.Error
		}
	} else {
		res = i.DB.Model(&ItemRecord{}).Where("name = ?", params[0]).Update(params[1], strings.Join(params[2:], " "))
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

func (i *Item) DeleteRecord(record string) (int64, error) {
	res := i.DB.Where("name = ?", record).Delete(ItemRecord{})
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
