package main

import "strings"

// CheckCommand checks if the command is valid and returns the appropriate response
func (i *Items) CheckCommand(cmd string) (string, error) {
	var msg string
	var params []string

	if len(strings.Split(cmd, " ")) > 2 {
		params = strings.Split(cmd, " ")[2:]
		cmd = strings.Join(strings.Split(cmd, " ")[:2], " ")
	}

	switch cmd {
	case "/start":
		msg = startMsg
	case "/list items":
		if len(params) == 0 {
			msg = i.ListItems("")
			if msg == "" {
				msg = noItems
			}
		} else {
			msg = i.ListItems(strings.Join(params[0:], " "))
		}
	case "/add item":
		msg, err := i.AddItem(params)
		if err != nil {
			return "", err
		}
		return msg, nil
	case "/update item":
		msg, err := i.UpdateItem(params)
		if err != nil {
			return "", err
		}
		return msg, nil
	case "/delete item":
		if len(params) == 0 {
			msg = deleteChoose
		} else {
			msg = i.DeleteItem(params)
		}
	default:
		msg = invalidMsg
	}
	return msg, nil
}
