package main

import (
	"bare/cmd"
	"bare/utils/osutil"
)

func main() {
	// Check for .bare folder
	osutil.MakeInitFolder()
	cmd.Execute()
}
