package cmd

import (
	"bare/styles"
	"bare/utils/parser"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about any bare on github",
	Run: func(cmd *cobra.Command, args []string) {
		InfoBare(args[0])
	},
}

func InfoBare(bareName string) {
	user, repo, branch := parser.ParseGithubRepo(bareName)
	parser.GetRecipe(user, repo, branch)
	// TODO Add variant description
	fmt.Println(styles.StatusPrompt.Render("Author   :"), parser.BareObj.Author)
	fmt.Println(styles.StatusPrompt.Render("Barename :"), parser.BareObj.BareName, parser.BareObj.Version)
	fmt.Println(styles.StatusPrompt.Render("Variants :"))
	for _, variant := range parser.BareObj.Variants {
		fmt.Println("*", variant)
	}
}
