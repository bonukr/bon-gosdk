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
	Name    string `json:"name"`    // 인터페이스 이름
	Type    string `json:"type"`    // 인터페이스 타입 (Physical, Virtual 등)
	Mac     string `json:"mac"`     // MAC 주소
	Cidr    string `json:"cidr"`    // Classless Inter-Domain Routing
	Ip      string `json:"ip"`      // IPv4 주소
	Netmask string `json:"netmask"` // 넷마스크
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
				fmt.Printf("_butils.GetPhysicalInterfaces: os.Stat failed: %s", err)
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
			fmt.Printf("_butils.GetPhysicalInterfaces: iface.Addrs failed: %s", err)
		} else {
			for _, addr := range addrs {
				// CIDR를 넷마스크로 변환
				ip, ipNet, err := net.ParseCIDR(addr.String())
				if err != nil {
					fmt.Printf("_butils.GetPhysicalInterfaces: net.ParseCIDR failed: %s", err)
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

type ResNetworkInterfaceAddress struct {
	Cidr    string `json:"cidr"`    // Classless Inter-Domain Routing
	Ip      string `json:"ip"`      // IPv4 주소
	Netmask string `json:"netmask"` // 넷마스크
}

type ResNetworkInterface struct {
	Type      string                       `json:"type"`                // 인터페이스 타입 (Physical, Virtual 등)
	Name      string                       `json:"name"`                // 인터페이스 이름
	Mac       string                       `json:"mac"`                 // MAC 주소
	Addresses []ResNetworkInterfaceAddress `json:"addresses,omitempty"` // Addresses
}

// GetPhysicalInterfaces : 물리 네트워크 인터페이스 목록 조회 함수
func GetNetworkInterfaces() (ret []ResNetworkInterface, err error) {
	// get interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %v", err)
	}

	// append to result
	for _, iface := range ifaces {
		// check - name
		if len(iface.Name) == 0 {
			continue
		}

		// append
		ret = append(ret, ResNetworkInterface{
			Type:      getNetworkType(&iface),
			Name:      iface.Name,
			Mac:       iface.HardwareAddr.String(),
			Addresses: getNetworkIpV4Address(&iface),
		})
	}

	// return
	return ret, nil
}

func getNetworkType(iface *net.Interface) string {
	const def_physical_path = "/sys/class/net/%s/device"
	const def_bonding_path = "/sys/class/net/%s/bonding/"
	const def_bridge_path = "/sys/class/net/%s/bridge/"
	const def_vlan_path = "/sys/class/net/%s/real_dev"

	// check
	if iface == nil {
		return "unknown"
	}

	// loopback
	if (iface.Flags & net.FlagLoopback) != 0 {
		return "loopback"
	}

	// physical
	if _, err := os.Stat(fmt.Sprintf(def_physical_path, iface.Name)); err == nil {
		return "physical"
	} else if !os.IsNotExist(err) {
		return "unknown"
	}

	// bonding
	if _, err := os.Stat(fmt.Sprintf(def_bonding_path, iface.Name)); err == nil {
		return "bonding"
	} else if !os.IsNotExist(err) {
		return "unknown"
	}

	// bridge
	if _, err := os.Stat(fmt.Sprintf(def_bridge_path, iface.Name)); err == nil {
		return "bridge"
	} else if !os.IsNotExist(err) {
		return "unknown"
	}

	// vlan
	if _, err := os.Stat(fmt.Sprintf(def_vlan_path, iface.Name)); err == nil {
		return "vlan"
	} else if !os.IsNotExist(err) {
		return "unknown"
	}

	// etc
	return "virtual"
}

func getNetworkIpV4Address(iface *net.Interface) (ret []ResNetworkInterfaceAddress) {
	// check
	if iface == nil {
		return
	}

	// addresses
	addrs, err := iface.Addrs()
	if err != nil {
		fmt.Printf("_butils.getNetworkIpV4Address: iface.Addrs failed: %s", err)
		return
	}

	// append to result
	for _, addr := range addrs {
		// CIDR를 넷마스크로 변환
		ip, ipNet, err := net.ParseCIDR(addr.String())
		if err != nil {
			fmt.Printf("_butils.getNetworkIpV4Address: net.ParseCIDR failed: %s", err)
			return
		}

		// IPv4만 처리 (IPv6은 제외)
		if ip.To4() == nil {
			continue
		}

		// append
		ret = append(ret, ResNetworkInterfaceAddress{
			Cidr:    addr.String(),
			Ip:      ip.String(),
			Netmask: net.IP(ipNet.Mask).String(),
		})
	}

	return
}
