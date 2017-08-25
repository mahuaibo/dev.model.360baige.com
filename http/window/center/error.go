package center

import (
//"errors"
)

var ErrorMap map[string]string

func init() {
	ErrorMap = map[string]string{
		"Normal":      "200",
		"ErrorSystem": "400",
		"ErrorLogic":  "500",
	}
}
