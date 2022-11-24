package iutils

import (
	"strings"
	"unicode"
)

// TODO move to friendly
func GetNameFromAddress(addr string) string {
	arr := strings.Split(strings.Split(addr, "@")[0], ".")
	res := ""
	for i := range arr {
		if len(arr[i]) == 1 {
			continue
		}
		res += string(append([]rune{unicode.ToUpper([]rune(arr[i])[0])}, []rune(arr[i])[1:]...))
	}
	return res
}
