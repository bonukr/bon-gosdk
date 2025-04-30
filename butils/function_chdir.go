package butils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ChangeWorkingDirToBinPath(isWinSvc bool) error {
	var ex string
	var err error

	if strings.EqualFold(runtime.GOOS, "windows") {
		if isWinSvc {
			ex, err = os.Executable()
			if err != nil {
				return fmt.Errorf("get image(exe) path failed: %s", err.Error())
			}
		} else {
			ex, err = filepath.Abs("./")
			if err != nil {
				return fmt.Errorf("get image(exe) path failed: %s", err.Error())
			}
			ex += "\\"
		}

	} else {
		ex, err = os.Executable()
		if err != nil {
			return fmt.Errorf("get image(exe) path failed: %s", err.Error())
		}
	}

	exPath := filepath.Dir(ex)
	err = os.Chdir(exPath)
	if err != nil {
		return fmt.Errorf("os.Chdir failed: %s", err.Error())
	}

	return nil
}
