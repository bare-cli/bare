package osutil

import (
	"log"
	"os"
	"path/filepath"
)

func GetBarePath() string {
	if os.Getenv("XDG_DATA_HOME") != "" {
		return os.Getenv("XDG_DATA_HOME")
	}
	return filepath.Join(os.Getenv("HOME"), ".local", "share")
}

func MakeInitFolder() {
	// Create Bare path
	barePath := GetBarePath()
	if !Exists(filepath.Join(barePath)) {
		if err := os.Mkdir(barePath, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

}

func MakeDownloadFolder() {
	barePath := GetBarePath()
	if !Exists(filepath.Join(barePath, "tmp")) {
		if err := os.Mkdir(filepath.Join(barePath, "tmp"), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
