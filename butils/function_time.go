package butils

import (
	"strings"
	"time"
)

var (
	_dft_time = time.Time{}
	_dft_loc_kst, _ = time.LoadLocation("Asia/Seoul")
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

// RFC3339 문자열 형식의 시간을 time 으로 변환
// 입력 : 
//		- in : "2026-02-19T01:23:45Z"
// 출력 : 
// 		- 실패 : time.Time 반환 (zero time)
// 		- 성공 : time.Time 반환
func Rfc3339StrToTime(in string) (time.Time) {
	out, err := time.Parse(time.RFC3339, in)
	if err != nil {
		return _dft_time
	} else {
		return out
	}
}

// time 을 KST 로 변환
// 출력 : 
// 		- 실패 : time.Time 반환 (zero time)
// 		- 성공 : time.Time 반환
func TimeToKST(in time.Time) time.Time {
	// Location 재할당
	if _dft_loc_kst == nil {
		_dft_loc_kst, _ = time.LoadLocation("Asia/Seoul")
		if _dft_loc_kst == nil {
			return _dft_time
		}
	}

	// 이미 Asia/Seoul이면 변환 없이 바로 리턴
	if in.Location() != nil && in.Location().String() == "Asia/Seoul" {
		return in
	}

	// KST 변환
	return in.In(_dft_loc_kst)
}
