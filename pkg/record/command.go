package record

import (
	"strings"
)

// CheckCommand checks if the command is valid and
// returns the appropriate response
func (r *RecordDB) CheckCommand(data string) (string, error) {
	cmd, params := checkParams(data)

	switch cmd {
	case "/start":
		return startMsg, nil
	case "/listitems":
		return r.List(params)
	case "/showitem":
		return r.Show(params)
	case "/additem":
		return r.Add(params)
	case "/updateitem":
		return r.Update(params)
	case "/deleteitem":
		return r.Delete(params)
	default:
		// add end case
		// check for previous command
		return invalidMsg, nil
	}
}

func checkParams(data string) (string, []string) {
	splitData := strings.Split(data, " ")
	return splitData[0], splitData[1:]
}
