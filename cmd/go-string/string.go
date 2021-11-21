package gostring

import (
	"fmt"
	"strings"
)

func Format(format string, mapArgs map[string]interface{}) string {
	if mapArgs == nil {
		return format
	}

	args := []string{}
	for key, value := range mapArgs {
		strValue := fmt.Sprint(value)
		curlyKey := "{" + key + "}"
		args = append(args, curlyKey, strValue)
	}
	r := strings.NewReplacer(args...)

	return r.Replace(format)
}
