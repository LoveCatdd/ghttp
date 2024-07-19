package utils

import "strings"

var buffer strings.Builder

func Concat(per, next, pattern string) string {
	buffer.Reset()
	buffer.WriteString(per)
	if pattern != "" {
		buffer.WriteString(pattern)
	}
	buffer.WriteString(next)
	return buffer.String()
}

func ToLower(str string) string {
	return strings.ToLower(str)
}
