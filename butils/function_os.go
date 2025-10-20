package butils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func IsLinux() bool {
	if runtime.GOOS == "linux" {
		return true
	} else {
		return false
	}
}

type ResLinuxOsRelease struct {
	Name string

	VersionString string
	Version       struct {
		Major int
		Minor int
		Patch int
		Extra string
	}
}

func GetLinuxOsRelease() (ResLinuxOsRelease, error) {
	const fname = "/etc/os-release"
	const keyname_name = "name"
	const keyname_version = "version_id"
	var ret ResLinuxOsRelease

	// check
	if !IsLinux() {
		return ret, fmt.Errorf("not linux")
	}

	// open
	file, err := os.Open(fname)
	if err != nil {
		return ret, err
	}
	defer file.Close()

	// scan
	datas := make(map[string]string)
	{
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			// skip empty lines and comments
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}

			// split key=value
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.ToLower(strings.TrimSpace(parts[0]))
			value := strings.Trim(strings.TrimSpace(parts[1]), `"`) // remove quotes if present
			datas[key] = value
		}

		if err := scanner.Err(); err != nil {
			return ret, err
		}
	}

	// result
	ret.Name = datas[keyname_name]
	ret.VersionString = datas[keyname_version]
	{
		// split
		splitData := strings.SplitN(ret.VersionString, "-", 2)

		// version (숫자.숫자.숫자 형태 추출)
		if len(splitData) >= 1 {
			tmp := strings.Split(splitData[0], ".")
			if len(tmp) >= 1 {
				ret.Version.Major, _ = strconv.Atoi(tmp[0])
			}
			if len(tmp) >= 2 {
				ret.Version.Minor, _ = strconv.Atoi(tmp[1])
			}
			if len(tmp) >= 3 {
				ret.Version.Patch, _ = strconv.Atoi(tmp[2])
			}
		}

		// extra
		if len(splitData) >= 2 {
			ret.Version.Extra = strings.TrimSpace(splitData[1])
		}
	}

	return ret, nil
}

type ResLinuxKernelVersion struct {
	VersionString string
	Version       struct {
		Major int
		Minor int
		Patch int
		Extra string
	}
}

func GetLinuxKernelVersion() (ResLinuxKernelVersion, error) {
	const execCmd = "/usr/bin/uname"
	const execArgs = "-r"
	var ret ResLinuxKernelVersion

	// check
	if !IsLinux() {
		return ret, fmt.Errorf("not linux")
	}

	// get
	var data string
	{
		var out bytes.Buffer
		cmd := exec.Command(execCmd, execArgs)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			return ret, err
		}

		data = strings.TrimSpace(out.String())
	}

	// result
	{
		// full
		ret.VersionString = strings.TrimSpace(data)

		// split
		splitData := strings.SplitN(data, "-", 2)

		// version (숫자.숫자.숫자 형태 추출)
		if len(splitData) >= 1 {
			tmp := strings.Split(splitData[0], ".")
			if len(tmp) >= 1 {
				ret.Version.Major, _ = strconv.Atoi(tmp[0])
			}
			if len(tmp) >= 2 {
				ret.Version.Minor, _ = strconv.Atoi(tmp[1])
			}
			if len(tmp) >= 3 {
				ret.Version.Patch, _ = strconv.Atoi(tmp[2])
			}
		}

		// extra
		if len(splitData) >= 2 {
			ret.Version.Extra = strings.TrimSpace(splitData[1])
		}
	}

	// result
	return ret, nil
}
