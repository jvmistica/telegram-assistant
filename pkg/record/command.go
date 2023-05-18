package record

import (
	"strings"
)

// defaultResponse contains default responses when parameters are not given
var defaultResponse map[string]string = map[string]string{
	"/start":      ResponseStart,
	"/showitem":   ResponseShow,
	"/additem":    ResponseAdd,
	"/updateitem": ResponseUpdate,
	"/deleteitem": ResponseDelete,
}

// CheckCommand checks if the command is valid and
// returns the appropriate response
func (r *RecordDB) CheckCommand(data string) (string, error) {
	if defaultResponse[data] != "" {
		return defaultResponse[data], nil
	}

	cmd, params := checkParams(data)

	switch cmd {
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
		return ResponseInvalid, nil
	}
}

func checkParams(data string) (string, []string) {
	splitData := strings.Split(data, " ")
	return splitData[0], splitData[1:]
}
