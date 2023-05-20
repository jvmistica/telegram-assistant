package record

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

const (
	defaultTimeFormat = "2006/01/02"
	defaultUnit       = "piece(s)"
	defaultCurrency   = "EUR"
	filterByName      = "name = ?"
	itemTag           = "<item>"
)

// RecordDB contains the database instance used for record transactions
type RecordDB struct {
	DB *gorm.DB
}

// Item is the model used for item-specific records
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

// Add inserts a new record into a table
func (r *RecordDB) Add(params []string) (string, error) {
	item := strings.Join(params, " ")
	record := Item{Name: item, Currency: "EUR"}
	err := r.DB.Create(&record)
	if err.Error != nil {
		return "", err.Error
	}

	msg := strings.ReplaceAll(ResponseSuccessAdd, itemTag, item)
	return msg, nil
}

// Show returns the details of a specific record
func (r *RecordDB) Show(params []string) (string, error) {
	rec := strings.Join(params, " ")
	var (
		item    Item
		details string
	)

	res := r.DB.Where(filterByName, rec).Find(&item)
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected == 0 {
		return strings.ReplaceAll(ResponseItemNotExist, itemTag, rec), nil
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

// ListRecords returns a list of records from the "item" table
func (r *RecordDB) ListRecords(cmd []string) (string, error) {
	if len(cmd) != 0 && len(cmd) < 3 {
		return ResponseInvalidList, nil
	}

	var items []Item
	var res *gorm.DB
	if len(cmd) == 0 {
		res = r.DB.Find(&items)
		if res.RowsAffected == 0 {
			return ResponseNoItems, nil
		}
	}

	if len(cmd) > 2 && strings.Join(cmd[:2], " ") == "sort by" {
		sort := ""
		if len(cmd) > 3 {
			sort = cmd[3]
		}

		res = r.sortList(cmd[2], sort, &items)
		if res.RowsAffected == 0 {
			return ResponseNoItems, nil
		}
	}

	if len(cmd) > 4 && strings.Join(cmd[:2], " ") == "filter by" {
		res = r.DB.Where(fmt.Sprintf("%s %s '%s'", cmd[2], cmd[3], strings.Join(cmd[4:], " "))).Find(&items)
		if res.RowsAffected == 0 {
			return ResponseNoMatchFilter, nil
		}
	}

	return r.getItemList(items), nil
}

// List returns a list of records from a table
func (r *RecordDB) List(params []string) (string, error) {
	var (
		msg string
		err error
	)

	if len(params) == 0 || params[0] == "" {
		msg, err = r.ListRecords([]string{})
		if err != nil {
			return "", err
		}
		return msg, nil
	}

	msg, err = r.ListRecords(params)
	if err != nil {
		return "", err
	}

	return msg, nil
}

// Update updates a specific record
func (r *RecordDB) Update(params []string) (string, error) {
	if len(params) < 3 {
		return "", errors.New(ResponseInvalidUpdate)
	}

	msg, err := r.UpdateRecord(params)
	if err != nil {
		return "", err
	}

	return msg, nil
}

// Delete deletes a specific record
func (r *RecordDB) Delete(params []string) (string, error) {
	rec := strings.Join(params, " ")
	res := r.DB.Where(filterByName, rec).Delete(Item{})
	if res.RowsAffected == 0 {
		return strings.ReplaceAll(ResponseItemNotExist, itemTag, rec), nil
	}

	return strings.ReplaceAll(ResponseSuccessDelete, itemTag, rec), nil
}

// Import imports a list of records into a table
func (r *RecordDB) Import(records [][]string) (string, error) {
	if _, err := r.ImportRecords(records); err != nil {
		return "", err
	}

	msg := "Successfully imported items."
	return msg, nil
}

// ImportRecords imports a list of records into the "item" table
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

// UpdateRecord updates a specific "item" record
func (r *RecordDB) UpdateRecord(params []string) (string, error) {
	var (
		res   *gorm.DB
		item  string
		field string
		value string
	)

	keys := []string{"description", "amount", "unit", "calories", "category", "price", "currency", "expiration"}
	for i, param := range params {
		for _, key := range keys {
			if param == key {
				item = strings.Join(params[:i], " ")
				field = param
				value = strings.Join(params[i+1:], " ")
			}

			if param == "amount" {
				field = param
				value = params[i+1]
				f, err := strconv.ParseFloat(params[i+1], 32)
				if err != nil {
					return "", err
				}

				res = r.DB.Model(&Item{}).Where(filterByName, item).Updates(Item{Amount: float32(f), Unit: defaultUnit})
				if res.Error != nil {
					return "", res.Error
				}
			}

			if param == "price" {
				field = param
				value = params[i+1]
				f, err := strconv.ParseFloat(params[i+1], 32)
				if err != nil {
					return "", err
				}

				res = r.DB.Model(&Item{}).Where(filterByName, item).Updates(Item{Price: float32(f), Currency: defaultCurrency})
				if res.Error != nil {
					return "", res.Error
				}
			}
		}
	}

	res = r.DB.Model(&Item{}).Where(filterByName, item).Update(field, value)
	if res.Error != nil {
		return "", res.Error
	}

	if res.RowsAffected == 0 {
		return strings.ReplaceAll(ResponseItemNotExist, itemTag, item), nil
	}

	return strings.ReplaceAll(strings.ReplaceAll(ResponseSuccessUpdate, itemTag, item), "<field>", field), nil
}

// sortList sorts a list of records
func (r *RecordDB) sortList(field, sort string, items *[]Item) *gorm.DB {
	var res *gorm.DB
	if sort == "asc" || sort == "desc" {
		res = r.DB.Order(fmt.Sprintf("%s %s", field, sort)).Find(items)
	} else {
		res = r.DB.Order(field).Find(&items)
	}

	return res
}

// getItemList formats a list of items
func (r *RecordDB) getItemList(items []Item) string {
	var itemsList string
	for _, item := range items {
		itemsList += item.Name + "\n"
	}

	return itemsList
}
