package utils

import (
	"path"
	"runtime"
)

func GetCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
