package osutil

import (
	"log"
	"os"
	"path/filepath"
)

func MakeInitFolder() {
	// Get home directory
	homePath := os.Getenv("HOME")
	if !Exists(filepath.Join(homePath, ".bare")) {
		if err := os.Mkdir(homePath+"/.bare", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

}

func MakeDownloadFolder() {
	homePath := os.Getenv("HOME")
	if !Exists(filepath.Join(homePath, ".bare", "tmp")) {
		if err := os.Mkdir(filepath.Join(homePath, ".bare", "tmp"), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
