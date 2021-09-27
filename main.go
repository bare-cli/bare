package main

import (
	"bare/cmd"
	"bare/utils"
)

func main() {
	// Check for .bare folder
	utils.MakeInitFolder()
	cmd.Execute()
}
