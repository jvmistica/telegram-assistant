package record

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type Record interface {
	AddRecord(record string) error
	ShowRecord(record string) (string, error)
	ListRecords(params []string) (string, error)
	UpdateRecord(params []string) (string, error)
	DeleteRecord(record string) (int64, error)
	ImportRecords(records [][]string) (string, error)
}

type RecordDB struct {
	DB *gorm.DB
}

func Add(r Record, params []string) (string, error) {
	if len(params) == 0 {
		return addChoose, nil
	}

	rec := strings.Join(params, " ")
	err := r.AddRecord(rec)
	if err != nil {
		return "", err
	}

	msg := strings.ReplaceAll(addSuccess, "<item>", rec)
	return msg, nil
}

func Show(r Record, params []string) (string, error) {
	if len(params) == 0 {
		return addChoose, nil
	}

	rec := strings.Join(params, " ")
	msg, err := r.ShowRecord(rec)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func List(r Record, params []string) (string, error) {
	var (
		msg string
		err error
	)

	if len(params) == 0 {
		msg, err = r.ListRecords([]string{})
		if err != nil {
			return "", err
		}
	}

	msg, err = r.ListRecords(params)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func Update(r Record, params []string) (string, error) {
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

func Delete(r Record, params []string) (string, error) {
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
		msg = strings.ReplaceAll(itemNotExist, "<item>", rec)
	} else {
		msg = strings.ReplaceAll(deleteSuccess, "<item>", rec)
	}

	return msg, nil
}

func Import(r Record, records [][]string) (string, error) {
	if _, err := r.ImportRecords(records); err != nil {
		return "", err
	}

	msg := "Successfully imported items."
	return msg, nil
}
