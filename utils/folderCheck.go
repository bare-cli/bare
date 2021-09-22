package utils

import "os"

func CheckFolder(folderPath string) bool {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return false
	}

	return true
}