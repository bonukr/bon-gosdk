package butils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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

func GetLinuxOsRelease() (string, string, error) {
	const fname = "/etc/os-release"

	// check
	if !IsLinux() {
		return "", "", fmt.Errorf("not linux")
	}

	// open
	file, err := os.Open(fname)
	if err != nil {
		return "", "", err
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

			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), `"`) // remove quotes if present
			datas[key] = value
		}

		if err := scanner.Err(); err != nil {
			return "", "", err
		}
	}

	// revise (lowercase to key)
	for k, v := range datas {
		datas[k] = strings.ToLower(v)
	}

	// result
	name := datas["name"]
	version := datas["version_id"]
	return name, version, nil
}

type ResLinuxKernelVersion struct {
	Version string
	Major   int
	Minor   int
	Patch   int
	Extra   string
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
		splitData := strings.SplitN(data, "-", 2)
		if len(splitData) == 0 {
			return ret, fmt.Errorf("invalid kernel version format: %s", data)
		}

		// version, extra
		ret.Version = strings.TrimSpace(splitData[0])
		if len(splitData) >= 2 {
			ret.Extra = strings.TrimSpace(splitData[1])
		}

		// major, minor, patch (숫자.숫자.숫자 형태 추출)
		{
			re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)`)
			matches := re.FindStringSubmatch(splitData[0])
			if len(matches) != 4 {
				return ret, fmt.Errorf("invalid kernel version format: %s", data)
			}

			ret.Major, _ = strconv.Atoi(matches[1])
			ret.Minor, _ = strconv.Atoi(matches[2])
			ret.Patch, _ = strconv.Atoi(matches[3])
		}
	}

	// result
	return ret, nil
}
