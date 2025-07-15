package butils

import (
	"strings"
)

func RemoveNewlines(s string) string {
	s = strings.ReplaceAll(s, "\r\n", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}
