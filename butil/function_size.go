package butil

import (
	"strconv"
	"strings"
)

// Kibibytes to Kilobytes
//   - Kilobytes = Kibibytes x 1024 ÷ 1000
//   - Kilobytes = Kibibytes x 1.024
func KiByteToKByte(kibibytes int) int {
	return int((float64(kibibytes) * 1.024) / 1)
}

func KiByteToKByteStr(kibibytes int) string {
	return strconv.Itoa(KiByteToKByte(kibibytes))
}

// Mebibytes to Megabytes
//   - Megabytes = Mebibytes x (1024x1024) ÷ (1000x1000)
//   - Megabytes = Mebibytes x 1.048576
func MiByteToMByte(mebibyte int) int {
	return int((float64(mebibyte) * 1.048576) / 1)
}

func MiByteToMByteStr(mebibyte int) string {
	return strconv.Itoa(MiByteToMByte(mebibyte))
}

// Gibibytes to Gigabytes
//   - Gigabytes = Gibibytes x (1024x1024x1024) ÷ (1000x1000x1000)
//   - Gigabytes = Gibibytes x 1.073741824
func GiByteToGByte(gibibyte int) int {
	return int((float64(gibibyte) * 1.073741824) / 1)
}

func GiByteToGByteStr(gibibyte int) string {
	return strconv.Itoa(GiByteToGByte(gibibyte))
}

// 함수 내에서 unit 을 ToLower + TrimSpace 처리함
// unit : kb, mb, gb, tb
func UnitByteToByte(val uint64, unit string) uint64 {
	unit = strings.ToLower(unit)
	unit = strings.TrimSpace(unit)
	const mul uint64 = 1024

	if unit == "kb" {
		return val * mul
	} else if unit == "mb" {
		return val * mul * mul
	} else if unit == "gb" {
		return val * mul * mul * mul
	} else if unit == "tb" {
		return val * mul * mul * mul * mul
	}

	return 0
}
