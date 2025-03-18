package butil

import "strconv"

// Kibibytes to Kilobytes
//	- Kilobytes = Kibibytes x 1024 ÷ 1000
//	- Kilobytes = Kibibytes x 1.024
func KiByteToKByte(kibibytes int) int {
	return int((float64(kibibytes) * 1.024) / 1)
}

func KiByteToKByteStr(kibibytes int) string {
	return strconv.Itoa(KiByteToKByte(kibibytes))
}

// Mebibytes to Megabytes
//	- Megabytes = Mebibytes x (1024x1024) ÷ (1000x1000)
//	- Megabytes = Mebibytes x 1.048576
func MiByteToMByte(mebibyte int) int {
	return int((float64(mebibyte) * 1.048576) / 1)
}

func MiByteToMByteStr(mebibyte int) string {
	return strconv.Itoa(MiByteToMByte(mebibyte))
}

// Gibibytes to Gigabytes
//	- Gigabytes = Gibibytes x (1024x1024x1024) ÷ (1000x1000x1000)
//	- Gigabytes = Gibibytes x 1.073741824
func GiByteToGByte(gibibyte int) int {
	return int((float64(gibibyte) * 1.073741824) / 1)
}

func GiByteToGByteStr(gibibyte int) string {
	return strconv.Itoa(GiByteToGByte(gibibyte))
}
