package butils

import "fmt"

func CidrToNetmask(bits int) (string, error) {
	if bits < 0 || bits > 32 {
		return "", fmt.Errorf("invalid bit count: %d. It must be between 0 and 32", bits)
	}

	// 비트를 서브넷 마스크로 변환
	mask := [4]byte{}
	for i := 0; i < 4; i++ {
		if bits >= 8 {
			mask[i] = 255
			bits -= 8
		} else {
			mask[i] = byte(255 << (8 - bits) & 255)
			break
		}
	}

	// 점-십진수 형식으로 변환
	netmask := fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])

	return netmask, nil
}
