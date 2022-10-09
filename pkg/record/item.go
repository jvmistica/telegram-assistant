package record

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	defaultTimeFormat = "2006/01/02"
	filterByName      = "name = ?"
	itemTag           = "<item>"
)

type Item struct {
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

func (r *RecordDB) ImportRecords(records [][]string) (string, error) {
	for _, row := range records {
		amount, err := strconv.ParseFloat(row[2], 32)
		if err != nil {
			return "", err
		}

		calories, err := strconv.Atoi(row[4])
		if err != nil {
			return "", err
		}

		price, err := strconv.ParseFloat(row[6], 32)
		if err != nil {
			return "", err
		}

		expiration, err := time.Parse(defaultTimeFormat, row[8])
		if err != nil {
			return "", err
		}

		rec := Item{Name: row[0], Description: row[1], Amount: float32(amount), Unit: row[3],
			Calories: uint16(calories), Category: row[5], Price: float32(price), Currency: row[7], Expiration: expiration}
		if err := r.DB.Create(&rec); err.Error != nil {
			return "", err.Error
		}
	}

	return "", nil
}

func (r *RecordDB) AddRecord(record string) error {
	rec := Item{Name: record}
	err := r.DB.Create(&rec)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (r *RecordDB) ShowRecord(record string) (string, error) {
	var (
		item    Item
		details string
	)

	res := r.DB.Where(filterByName, record).Find(&item)
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected == 0 {
		return strings.ReplaceAll(itemNotExist, itemTag, record), nil
	}

	category := "_Uncategorized_"
	if item.Category != "" {
		category = item.Category
	}

	description := "_No description_"
	if item.Description != "" {
		description = item.Description
	}

	exp := item.Expiration
	expiration := exp.Format(defaultTimeFormat)
	if expiration[0] == '0' {
		expiration = "_Not set_"
	}

	details = fmt.Sprintf("*%s* (%s)\n\n%s\nAmount: %.2f %s\nPrice: %.2f %s\nExpiration: %s",
		item.Name, category, description, item.Amount, item.Unit, item.Price, item.Currency, expiration)

	return details, nil
}

func (r *RecordDB) sortList(field, sort string, items *[]Item) *gorm.DB {
	var res *gorm.DB
	if sort == "asc" || sort == "desc" {
		res = r.DB.Order(fmt.Sprintf("%s %s", field, sort)).Find(items)
	} else {
		res = r.DB.Order(field).Find(&items)
	}

	return res
}

func (r *RecordDB) ListRecords(cmd []string) (string, error) {
	var (
		items     []Item
		itemsList string
		res       *gorm.DB
	)

	if len(cmd) == 0 {
		res = r.DB.Find(&items)
		if res.RowsAffected == 0 {
			return noItems, nil
		}
	}

	if len(cmd) == 4 && strings.Join(cmd[:2], " ") == "sort by" {
		res = r.sortList(cmd[2], cmd[3], &items)
		if res.RowsAffected == 0 {
			return noItems, nil
		}
	}

	if len(cmd) == 5 && strings.Join(cmd[:2], " ") == "filter by" {
		res = r.DB.Where(fmt.Sprintf("%s %s '%s'", cmd[2], cmd[3], strings.Join(cmd[4:], " "))).Find(&items)
		if res.RowsAffected == 0 {
			return noMatchFilter, nil
		}
	}

	if len(cmd) != 0 && len(cmd) != 4 && len(cmd) != 6 {
		return invalidListMsg, nil
	}

	for _, item := range items {
		itemsList += item.Name + "\n"
	}

	if res != nil && res.Error != nil {
		return "", res.Error
	}

	return itemsList, nil
}

func (r *RecordDB) UpdateRecord(params []string) (string, error) {
	var (
		msg string
		res *gorm.DB
	)

	if params[1] == "amount" && len(params) > 3 {
		f, err := strconv.ParseFloat(params[2], 32)
		if err != nil {
			return "", err
		}

		res = r.DB.Model(&Item{}).Where(filterByName, params[0]).Updates(Item{Amount: float32(f), Unit: params[3]})
		if res.Error != nil {
			return "", res.Error
		}
	} else if params[1] == "price" && len(params) > 3 {
		f, err := strconv.ParseFloat(params[2], 32)
		if err != nil {
			return "", err
		}

		res = r.DB.Model(&Item{}).Where(filterByName, params[0]).Updates(Item{Price: float32(f), Currency: params[3]})
		if res.Error != nil {
			return "", res.Error
		}
	} else {
		res = r.DB.Model(&Item{}).Where(filterByName, params[0]).Update(params[1], strings.Join(params[2:], " "))
		if res.Error != nil {
			return "", res.Error
		}
	}

	if res.RowsAffected == 0 {
		msg = strings.ReplaceAll(itemNotExist, itemTag, params[0])
	} else {
		msg = strings.ReplaceAll(strings.ReplaceAll(updateSuccess, itemTag, params[0]), "<field>", params[1])
	}

	return msg, nil
}

func (r *RecordDB) DeleteRecord(record string) (int64, error) {
	res := r.DB.Where(filterByName, record).Delete(Item{})
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}
