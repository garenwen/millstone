package utils

import (
	"os"
	"path/filepath"
)

const (
	LockKeyPrefix = "millstone_lock_"
)

func GetLockName(name, spec string) string {
	return LockKeyPrefix + filepath.Base(os.Args[0]) + name + spec
}
