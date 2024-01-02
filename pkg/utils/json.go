package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJsonPretty(data any) {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(bytes))
}

func AsJson(data any) []byte {
	bytes, _ := json.Marshal(data)
	return bytes
}
