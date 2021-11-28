package cmd

import (
	"bare/styles"
	"bare/utils/git"
	"bare/utils/osutil"
	"bare/utils/parser"
	"bare/utils/template"
	"bare/utils/ui"
	"bare/utils/validate"
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
		err := validate.ValidateArgCount(2, len(args))
		if err != nil {
			fmt.Println(styles.StatusError.Render("X"), err)
			os.Exit(1)
		}
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

	// Prompt project name and template
	TempObj.Name = ui.PromptString("Enter project name", parser.BareObj.BareName)
	TempObj.Template = ui.PromptSelect("Select template", parser.BareObj.Variants)

	// Prompt variables
	TempObj.Variables = make(map[string]string)
	for k, e := range parser.BareObj.Variables {
		TempObj.Variables[k] = ui.PromptString(k, e)
	}
	fmt.Println(TempObj)
	osutil.MakeDownloadFolder()
	err := git.DownloadZip(user, repo, branch, parser.BareObj.BareName)
	if err != nil {
		log.Fatal(err)
	}

	downloadZipName := parser.BareObj.BareName + ".zip"
	downloadZipPath := filepath.Join(BarePath, "tmp", downloadZipName)
	extractZipName := user + "_" + repo + "_" + branch
	extractZipPath := BarePath
	destiPath := filepath.Join(extractZipPath, extractZipName)

	// Create extracted folder at BarePath
	err = osutil.CreateIfNotExists(filepath.Join(extractZipPath, extractZipName), 0755)
	if err != nil {
		log.Fatal(err)
	}
	// Extract here
	osutil.Unzip(downloadZipPath, destiPath)

	// Execution
	// Variant path
	varPath := filepath.Join(BarePath, extractZipName, TempObj.Template)
	currDir, err := os.Getwd()
	targetPath := filepath.Join(currDir, desti)
	err = template.Execute(varPath, targetPath, TempObj.Variables)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(varPath)
}
