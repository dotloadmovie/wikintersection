package utils

import (
	"strings"
)

func MatchArray(arr []string, str string) bool {
	output := false

	for _, item := range arr {
		if strings.Contains(strings.ToLower(str), strings.ToLower(item)) {
			output = true
		}
	}

	return output
}
