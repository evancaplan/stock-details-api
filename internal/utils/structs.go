package utils

import "encoding/json"

func ToString(structz interface{}) string {
	out, err := json.Marshal(structz)
	if err != nil {
		panic(err)
	}

	return string(out)
}
