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
	Name    string
	Version string

	Major int
	Minor int
	Patch int
	Extra string
}

func GetLinuxOsRelease() (ResLinuxOsRelease, error) {
	const fname = "/etc/os-release"
	const keyname_name = "name"
	const keyname_version = "version_id"

	// init
	var ret ResLinuxOsRelease
	{
		ret.Major = -1
		ret.Minor = -1
		ret.Patch = -1
	}

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
	ret.Version = datas[keyname_version]
	{
		// split
		splitData := strings.SplitN(ret.Version, "-", 2)

		// version (숫자.숫자.숫자 형태 추출)
		if len(splitData) >= 1 {
			tmp := strings.Split(splitData[0], ".")
			if len(tmp) >= 1 {
				ret.Major, _ = strconv.Atoi(tmp[0])
			}
			if len(tmp) >= 2 {
				ret.Minor, _ = strconv.Atoi(tmp[1])
			}
			if len(tmp) >= 3 {
				ret.Patch, _ = strconv.Atoi(tmp[2])
			}
		}

		// extra
		if len(splitData) >= 2 {
			ret.Extra = strings.TrimSpace(splitData[1])
		}
	}

	return ret, nil
}

type ResLinuxKernelVersion struct {
	Version string

	Major int
	Minor int
	Patch int
	Extra string
}

func GetLinuxKernelVersion() (ResLinuxKernelVersion, error) {
	const execCmd = "/usr/bin/uname"
	const execArgs = "-r"

	// init
	var ret ResLinuxKernelVersion
	{
		ret.Major = -1
		ret.Minor = -1
		ret.Patch = -1
	}

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
		ret.Version = strings.TrimSpace(data)

		// split
		splitData := strings.SplitN(data, "-", 2)

		// version (숫자.숫자.숫자 형태 추출)
		if len(splitData) >= 1 {
			tmp := strings.Split(splitData[0], ".")
			if len(tmp) >= 1 {
				ret.Major, _ = strconv.Atoi(tmp[0])
			}
			if len(tmp) >= 2 {
				ret.Minor, _ = strconv.Atoi(tmp[1])
			}
			if len(tmp) >= 3 {
				ret.Patch, _ = strconv.Atoi(tmp[2])
			}
		}

		// extra
		if len(splitData) >= 2 {
			ret.Extra = strings.TrimSpace(splitData[1])
		}
	}

	// result
	return ret, nil
}
