package butils

import (
	"fmt"
	"net"
	"os"
)

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

type ResPhysicalInterface struct {
	Name    string // 인터페이스 이름
	Type    string // 인터페이스 타입 (Physical, Virtual 등)
	Mac     string // MAC 주소
	Cidr    string // Classless Inter-Domain Routing
	Ip      string // IPv4 주소
	Netmask string // 넷마스크
}

// GetPhysicalInterfaces : 물리 네트워크 인터페이스 목록 조회 함수
func GetPhysicalInterfaces() ([]ResPhysicalInterface, error) {
	const def_physical_path = "/sys/class/net/%s/device"
	var result []ResPhysicalInterface

	// get interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %v", err)
	}

	// iterate interfaces
	for _, iface := range ifaces {
		// check - loopback
		if (iface.Flags & net.FlagLoopback) != 0 {
			continue
		}
		// check - name
		if len(iface.Name) == 0 {
			continue
		}

		// filter - physical
		chkPath := fmt.Sprintf(def_physical_path, iface.Name)
		if _, err := os.Stat(chkPath); err != nil {
			if os.IsNotExist(err) {
				continue // device 디렉토리가 없으면 가상 인터페이스
			} else {
				fmt.Printf("_butils.GetPhysicalInterfaces: failed to os.Stat: %s", err)
				continue
			}
		}

		// new
		newIf := ResPhysicalInterface{
			Name: iface.Name,
			Type: "Physical",
		}

		// address
		if addrs, err := iface.Addrs(); err != nil {
			fmt.Printf("_butils.GetPhysicalInterfaces: failed to iface.Addrs: %s", err)
		} else {
			for _, addr := range addrs {
				// CIDR를 넷마스크로 변환
				ip, ipNet, err := net.ParseCIDR(addr.String())
				if err != nil {
					fmt.Printf("_butils.GetPhysicalInterfaces: failed to net.ParseCIDR: %s", err)
					continue
				}

				// IPv4만 처리 (IPv6은 제외)
				if ip.To4() == nil {
					continue
				}

				// result
				newIf.Mac = iface.HardwareAddr.String()
				newIf.Ip = ip.String()
				newIf.Cidr = addr.String()
				newIf.Netmask = net.IP(ipNet.Mask).String()
			}
		}

		// append to result
		result = append(result, newIf)
	}

	return result, nil
}
