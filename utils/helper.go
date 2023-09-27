package utils

import "strconv"

func StringToInt(from string) int {
	result, _ := strconv.Atoi(from)
	return result
}

func IntToString(from int) string {
	return strconv.Itoa(from)
}

func BoolToString(from bool) string {
	return strconv.FormatBool(from)
}

func IsEmptyString(str string) bool {
	return str == ""
}