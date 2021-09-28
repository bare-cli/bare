package utils

import (
	"log"
	"os"
)

func MakeInitFolder() {
	// Get home directory
	homePath := os.Getenv("HOME")
	if !Exists(homePath + "/.bare") {
		if err := os.Mkdir(homePath+"/.bare", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
