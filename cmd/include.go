package cmd

import (
	"bare/parser"
	"bare/styles"
	"bare/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(includeCmd)
}

var includeCmd = &cobra.Command{
	Use:   "include",
	Short: "include files to recipe.json `Include` field",
	Run: func(cmd *cobra.Command, args []string) {
		includeFiles(args)
	},
}

func includeFiles(objects []string) {
	parser.Parser("./recipe.json")
	incMap := make(map[string]bool)
	currDir, _ := os.Getwd()
	for _, objs := range parser.BareObj.Include {
		incMap[objs] = true
	}

	for _, objs := range objects {
		if incMap[objs] {
			continue
		} else {
			if utils.Exists(filepath.Join(currDir, objs)) {
				parser.BareObj.Include = append(parser.BareObj.Include, objs)
				fmt.Println(styles.InitSuccess.Render("[Add] " + objs))
			} else {
				fmt.Println(styles.InitError.Render("[Not found] "), objs)
			}
		}
	}
	parser.UpdateRecipe()
}
