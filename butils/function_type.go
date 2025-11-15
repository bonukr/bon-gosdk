package butils

import (
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func ToBool(in interface{}) (out bool) {
	switch v := in.(type) {
	case int:
		if v != 0 {
			out = true
		} else {
			out = false
		}
	case int64:
		if v != 0 {
			out = true
		} else {
			out = false
		}
	case uint:
		if v != 0 {
			out = true
		} else {
			out = false
		}
	case uint64:
		if v != 0 {
			out = true
		} else {
			out = false
		}
	case uint32:
		if v != 0 {
			out = true
		} else {
			out = false
		}
	case float64:
		if v != 0 {
			out = true
		} else {
			out = false
		}
	case float32:
		if v != 0 {
			out = true
		} else {
			out = false
		}
	case string:
		if strings.EqualFold(v, "true") || strings.EqualFold(v, "t") || strings.EqualFold(v, "yes") || strings.EqualFold(v, "y") || strings.EqualFold(v, "1") || strings.EqualFold(v, "on") {
			out = true
		} else {
			out = false
		}
	case []byte:
		tmp := string(v)
		if strings.EqualFold(tmp, "true") || strings.EqualFold(tmp, "t") || strings.EqualFold(tmp, "yes") || strings.EqualFold(tmp, "y") || strings.EqualFold(tmp, "1") || strings.EqualFold(tmp, "on") {
			out = true
		} else {
			out = false
		}
	case bool:
		if in.(bool) {
			out = true
		} else {
			out = false
		}
	default:
		out = false
	}

	return
}

func ToInt(in interface{}) (out int) {
	var err error

	switch v := in.(type) {
	case int:
		out = v
	case int64:
		out = int(v)
	case uint:
		out = int(v)
	case uint64:
		out = int(v)
	case uint32:
		out = int(v)
	case float64:
		out = int(v)
	case float32:
		out = int(v)
	case byte:
		out = int(v)
	case string:
		out, err = strconv.Atoi(string(v))
		if err != nil {
			tmp, _ := strconv.ParseFloat(v, int(unsafe.Sizeof(int(0)))*8)
			out = int(tmp)
		}
	case []byte:
		out, err = strconv.Atoi(string(v))
		if err != nil {
			tmp, _ := strconv.ParseFloat(string(v), int(unsafe.Sizeof(int(0)))*8)
			out = int(tmp)
		}
	case bool:
		if in.(bool) {
			out = 1
		} else {
			out = 0
		}
	default:
		out = 0
	}

	return
}

func ToInt64(in interface{}) (out int64) {
	var err error

	switch v := in.(type) {
	case int:
		out = int64(v)
	case int64:
		out = int64(v)
	case uint:
		out = int64(v)
	case uint64:
		out = int64(v)
	case uint32:
		out = int64(v)
	case float64:
		out = int64(v)
	case float32:
		out = int64(v)
	case byte:
		out = int64(v)
	case string:
		out, err = strconv.ParseInt(v, 10, 0)
		if err != nil {
			tmp, _ := strconv.ParseFloat(v, int(unsafe.Sizeof(int(0)))*8)
			out = int64(tmp)
		}
	case []byte:
		out, err = strconv.ParseInt(string(v), 10, 0)
		if err != nil {
			tmp, _ := strconv.ParseFloat(string(v), int(unsafe.Sizeof(int(0)))*8)
			out = int64(tmp)
		}
	case bool:
		if in.(bool) {
			out = 1
		} else {
			out = 0
		}
	default:
		out = 0
	}

	return
}

func ToUint(in interface{}) (out uint) {
	switch v := in.(type) {
	case int:
		out = uint(v)
	case int64:
		out = uint(v)
	case uint:
		out = uint(v)
	case uint64:
		out = uint(v)
	case uint32:
		out = uint(v)
	case float64:
		out = uint(v)
	case float32:
		out = uint(v)
	case string:
		tmp, _ := strconv.ParseUint(v, 10, 0)
		out = uint(tmp)
	case byte:
		out = uint(v)
	case []byte:
		tmp, _ := strconv.ParseUint(string(v), 10, 0)
		out = uint(tmp)
	case bool:
		if in.(bool) {
			out = 1
		} else {
			out = 0
		}
	default:
		out = 0
	}

	return
}

func ToUint64(in interface{}) (out uint64) {
	var err error

	switch v := in.(type) {
	case int:
		out = uint64(v)
	case int64:
		out = uint64(v)
	case uint:
		out = uint64(v)
	case uint64:
		out = v
	case uint32:
		out = uint64(v)
	case float64:
		out = uint64(v)
	case float32:
		out = uint64(v)
	case byte:
		out = uint64(v)
	case string:
		out, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			tmp, _ := strconv.ParseFloat(v, int(unsafe.Sizeof(int(0)))*8)
			out = uint64(tmp)
		}
	case []byte:
		out, err = strconv.ParseUint(string(v), 10, 64)
		if err != nil {
			tmp, _ := strconv.ParseFloat(string(v), int(unsafe.Sizeof(int(0)))*8)
			out = uint64(tmp)
		}
	case bool:
		if in.(bool) {
			out = 1
		} else {
			out = 0
		}
	default:
		out = 0
	}

	return
}

func ToFloat64(in interface{}) (out float64) {
	switch v := in.(type) {
	case int:
		out = float64(v)
	case int64:
		out = float64(v)
	case uint:
		out = float64(v)
	case uint64:
		out = float64(v)
	case uint32:
		out = float64(v)
	case float64:
		out = float64(v)
	case float32:
		out = float64(v)
	case string:
		out, _ = strconv.ParseFloat(in.(string), 64)
	case []byte:
		out, _ = strconv.ParseFloat(string(v), 64)
	case bool:
		if in.(bool) {
			out = 1
		} else {
			out = 0
		}
	default:
		out = 0
	}

	return
}

func ToString(in interface{}) (out string) {
	switch v := in.(type) {
	case float64:
		out = strconv.FormatFloat(in.(float64), 'f', 6, 64)
	case float32:
		out = strconv.FormatFloat(float64(in.(float32)), 'f', 6, 32)
	case int:
		out = strconv.FormatInt(int64(in.(int)), 10)
	case int64:
		out = strconv.FormatInt(in.(int64), 10)
	case uint:
		out = strconv.FormatUint(uint64(in.(uint)), 10)
	case uint64:
		out = strconv.FormatUint(in.(uint64), 10)
	case uint32:
		out = strconv.FormatUint(uint64(in.(uint32)), 10)
	case string:
		out = in.(string)
	case *string:
		if v == nil {
			return ""
		} else {
			out = *v
		}
	case []byte:
		out = string(v)
	case bool:
		if in.(bool) {
			out = "true"
		} else {
			out = "false"
		}
	case time.Time:
		out = v.Format(time.RFC3339)
	case *time.Time:
		if v == nil {
			return ""
		} else {
			out = v.Format(time.RFC3339)
		}
	default:
		out = ""
	}

	return
}

func ToMaskString(in interface{}) (out string) {
	tmp := ToString(in)
	out = strings.Repeat("*", len(tmp))

	return
}

// in: 입력값 (ex 0xA1b2, 0Xa1B2)
//   - 빈문자열(") 일 경우 dft 반환됨.
//   - 변환 실패 또는 int64 최대 값을 초과한 경우 dft 값이 반환됨
func HexToInt64(in string, dft int64) int64 {
	// Trim spaces
	in = strings.TrimSpace(in)

	// Remove "0x" or "0X" prefix if exists
	if strings.HasPrefix(in, "0x") || strings.HasPrefix(in, "0X") {
		in = in[2:]
	}

	// Validate: empty string is not allowed
	if in == "" {
		return dft
	}

	// Parse using 16 base
	out, err := strconv.ParseInt(in, 16, 64)
	if err != nil {
		return dft
	}

	return out
}
