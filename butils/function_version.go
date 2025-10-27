package butils

import (
	"strconv"
	"strings"
)

func ParseVersion(ver string) (major int, minor int, patch int) {
	var err error

	parts := strings.SplitN(ver, ".", 3)
	tmp := len(parts)
	if tmp >= 1 {
		if major, err = strconv.Atoi(strings.TrimSpace(parts[0])); err != nil {
			major = -1
		}
	}
	if tmp >= 2 {
		if major, err = strconv.Atoi(strings.TrimSpace(parts[1])); err != nil {
			minor = -1
		}
	}
	if tmp >= 3 {
		if major, err = strconv.Atoi(strings.TrimSpace(parts[2])); err != nil {
			patch = -1
		}
	}

	return
}
