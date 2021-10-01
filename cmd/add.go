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
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add you current bare to repo",
	Long:  "same as short for addCmd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.InitStyle.Render("Bare Add"))
		addBare()
	},
}

func addBare() {
	recipePath := "./recipe.json"
	parser.Parser(recipePath)
	barePath := parser.BareObj.BarePath
	for _, objPath := range parser.BareObj.Include {
		sourcePath := filepath.Join(".", objPath)
		destiPath := filepath.Join(barePath, objPath)
		err := utils.CopyFileDirectory(sourcePath, destiPath)
		if err != nil {
			fmt.Print(styles.InitError.Render("[Error] "), styles.AddFileStlyle.Render(objPath))
			fmt.Println("")
			os.Exit(1)
		} else {
			fmt.Print(styles.InitSuccess.Render("[Success] "), styles.AddFileStlyle.Render(objPath))
			fmt.Println("")
		}
	}

}
