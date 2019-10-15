package utils

import (
	"encoding/json"
	"log"
)

func ParseError(err error) int {
	switch err.Error() {
	case "NOT_FOUND":
		return 404
	default:
		return 500
	}
}

func ApiResult(i interface{}) string {
	result, err := json.Marshal(i)

	if err != nil {
		log.Fatal("Error in convert to JSON")
		return ""
	}

	return string(result)
}