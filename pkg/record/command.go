package record

import "strings"

// CheckCommand checks if the command is valid and returns the appropriate response
func (i *Item) CheckCommand(cmd string) (string, error) {
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
		msg, err = List(i, params)
	case "/showitem":
		msg, err = Show(i, params)
	case "/additem":
		msg, err = Add(i, params)
	case "/updateitem":
		msg, err = Update(i, params)
	case "/deleteitem":
		msg, err = Delete(i, params)
	default:
		msg = invalidMsg
	}

	if err != nil {
		return "", err
	}

	return msg, nil
}
