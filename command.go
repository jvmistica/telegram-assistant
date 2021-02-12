package main

import "strings"

// CheckCommand checks if the command is valid and returns the appropriate response
func (i *Items) CheckCommand(cmd string) (string, error) {
	var (
		msg    string
		err    error
		params []string
	)

	if len(strings.Split(cmd, " ")) > 1 {
		params = strings.Split(cmd, " ")[1:]
		cmd = strings.Split(cmd, " ")[0]
	}

	switch cmd {
	case "/start":
		msg = startMsg
	case "/listitems":
		if len(params) == 0 {
			msg = i.ListItems("")
		} else {
			msg = i.ListItems(strings.Join(params[0:], " "))
		}
	case "/showitem":
		msg, err = i.ShowItem(params)
	case "/additem":
		msg, err = i.AddItem(params)
	case "/updateitem":
		msg, err = i.UpdateItem(params)
	case "/deleteitem":
		msg, err = i.DeleteItem(params)
	default:
		msg = invalidMsg
	}

	if err != nil {
		return "", err
	}

	return msg, nil
}
