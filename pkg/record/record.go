package record

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

// Record is an interface that can be used for
// performing record transactions
type Record interface {
	AddRecord(record string) error
	ShowRecord(record string) (string, error)
	ListRecords(params []string) (string, error)
	UpdateRecord(params []string) (string, error)
	DeleteRecord(record string) (int64, error)
	ImportRecords(records [][]string) (string, error)
}

// RecordDB contains the database instance
// used in performing record transactions
type RecordDB struct {
	DB *gorm.DB
}

// Add inserts a new record into a table
func (r *RecordDB) Add(params []string) (string, error) {
	if len(params) == 0 || params[0] == "" {
		return addChoose, nil
	}

	rec := strings.Join(params, " ")
	err := r.AddRecord(rec)
	if err != nil {
		return "", err
	}

	msg := strings.ReplaceAll(addSuccess, itemTag, rec)
	return msg, nil
}

// Show returns the details of a specific record
func (r *RecordDB) Show(params []string) (string, error) {
	if len(params) == 0 || params[0] == "" {
		return showChoose, nil
	}

	rec := strings.Join(params, " ")
	msg, err := r.ShowRecord(rec)
	if err != nil {
		return "", err
	}

	return msg, nil
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
	if len(params) == 0 {
		return updateChoose, nil
	}

	if len(params) < 3 {
		return "", errors.New(updateInvalid)
	}

	msg, err := r.UpdateRecord(params)
	if err != nil {
		return "", err
	}

	return msg, nil
}

// Delete deletes a specific record
func (r *RecordDB) Delete(params []string) (string, error) {
	var msg string

	if len(params) == 0 {
		return deleteChoose, nil
	}

	rec := strings.Join(params, " ")
	rows, err := r.DeleteRecord(rec)
	if err != nil {
		return "", err
	}

	if rows == 0 {
		msg = strings.ReplaceAll(itemNotExist, itemTag, rec)
	} else {
		msg = strings.ReplaceAll(deleteSuccess, itemTag, rec)
	}

	return msg, nil
}

// Import imports a list of records into a table
func (r *RecordDB) Import(records [][]string) (string, error) {
	if _, err := r.ImportRecords(records); err != nil {
		return "", err
	}

	msg := "Successfully imported items."
	return msg, nil
}
