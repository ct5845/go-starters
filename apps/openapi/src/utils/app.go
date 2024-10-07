package utils

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

var appDirOnce sync.Once
var appDir string

func AppDir() string {
	appDirOnce.Do(func() {
		exePath, err := os.Executable()

		if err != nil {
			log.Fatal(err)
		}

		appDir = filepath.Dir(exePath)
	})

	return appDir
}
