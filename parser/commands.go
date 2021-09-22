package parser

import "fmt"

func RecipeCmds(cmdArgs []string) {
	switch cmdArgs[0] {
	case "BARE":
		cmdBare(cmdArgs[1]);
	}
}

func cmdBare(bareName string) {
	fmt.Println(bareName)
}