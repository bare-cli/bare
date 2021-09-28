package cmd

import (
	"bare/parser"
	"bare/styles"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(includeCmd)
}

var includeCmd = &cobra.Command{
	Use : "include",
	Short: "include files to recipe.json `Include` field",
	Run: func(cmd *cobra.Command, args []string) {
		includeFiles(args)	
	},
}

func includeFiles(objects []string) {
	parser.Parser("./recipe.json")
	incMap := make(map[string]bool)
	for _, objs := range parser.BareObj.Include {
		incMap[objs] = true
	}

	for _, objs := range objects {
		if incMap[objs]{
			continue
		}else{
			parser.BareObj.Include = append(parser.BareObj.Include, objs)
			fmt.Println(styles.InitSuccess.Render("[Add] " + objs))
		}
	}
	parser.UpdateRecipe()
}

