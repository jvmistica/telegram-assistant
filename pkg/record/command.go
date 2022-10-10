package record

import (
	"strings"
)

// CheckCommand checks if the command is valid and
// returns the appropriate response
func (r *RecordDB) CheckCommand(cmd string) (string, error) {
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
		msg, err = r.List(params)
	case "/showitem":
		msg, err = r.Show(params)
	case "/additem":
		msg, err = r.Add(params)
	case "/updateitem":
		msg, err = r.Update(params)
	case "/deleteitem":
		msg, err = r.Delete(params)
	default:
		msg = invalidMsg
	}

	if err != nil {
		return "", err
	}

	return msg, nil
}
