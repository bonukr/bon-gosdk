package butils

import (
	"strings"
	"time"
)

func ShortDuration(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}

// name: "loc", "UTC", "America/New_York", "Asia/Seoul"
func SetDefaultTimeZone(name string) {
	//time.Local = time.FixedZone("UTC", 0)
	time.Local, _ = time.LoadLocation(name)
}
