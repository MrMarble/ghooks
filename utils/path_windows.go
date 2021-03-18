// +build windows

package utils

import (
	"path/filepath"
	"runtime"
)

func Absolute(path string) string {
	_, base, _, _ := runtime.Caller(0)
	return filepath.VolumeName(base) + string(filepath.Separator) + path
}
