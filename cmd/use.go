package cmd

import (
	"bare/utils/git"
	"bare/utils/osutil"
	"bare/utils/parser"
	"bare/utils/template"
	"bare/utils/ui"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(useCmd)
}

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Create a project from all your bares",
	Run: func(cmd *cobra.Command, args []string) {
		useBare(args[0], args[1])
		// bare use <bare-name> <destination>
	},
}

type NewTemplate struct {
	Name      string
	Template  string
	Variables map[string]string
	BarePath  string
}

var TempObj NewTemplate

func useBare(bareName, desti string) {
	user, repo, branch := parser.ParseGithubRepo(bareName)
	parser.GetRecipe(user, repo, branch)
	//ui.PromptStringNew("Enter project name", parser.BareObj.BareName)
	TempObj.Name = ui.PromptString("Enter project name", parser.BareObj.BareName)
	TempObj.Template = ui.PromptSelect("Select template", parser.BareObj.Variants)

	// Prompt variables
	TempObj.Variables = make(map[string]string)
	for k, e := range parser.BareObj.Variables {
		varName := ui.PromptString(k, e)
		TempObj.Variables[k] = varName
	}

	osutil.MakeDownloadFolder()
	err := git.DownloadZip(user, repo, branch, parser.BareObj.BareName)
	if err != nil {
		log.Fatal(err)
	}

	var downloadZipName string = parser.BareObj.BareName + ".zip"
	var downloadZipPath string = filepath.Join(BarePath, "tmp", downloadZipName)
	//var extractZipName string = repo + "-" + branch
	//var extractZipPath string = BarePath
	cwd, _ := os.Getwd()
	var destiPath string = filepath.Join(cwd, desti)
	// Extract here
	osutil.Unzip(downloadZipPath, destiPath)
	//
	// Execution
	// Variant path
	var varPath string = filepath.Join(BarePath, parser.BareObj.BareName, TempObj.Template)
	err = template.Execute(varPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(varPath)
}
