package cmd

import (
	"bare/utils/git"
	"bare/utils/osutil"
	"bare/utils/parser"
	"bare/utils/ui"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/walle/targz"
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
	// TODO
	//	[x] Parse github repo
	// 	[x] Get recipe.json
	// 	[x] prompt preferences
	//		[x] Prompt Project name
	//		[x] Prompt template
	//		[x] Prompt variables
	// 	[ ] If successful download .tar.gz
	// 	[ ] extract .tar.gz
	// 		[ ] If -C flag is present cache it.
	// 	[ ] Execute the required template
	// - Download the repo as .tar.gz

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

	var downloadZipName string = parser.BareObj.BareName + ".tar.gz"
	var extractZipName string = repo + "-" + branch
	err = targz.Extract(filepath.Join(BarePath, "tmp", downloadZipName), BarePath)
	if err != nil {
		log.Fatal("Error extracting template")
	}

	fmt.Println(downloadZipName, extractZipName)
	err = os.Rename(filepath.Join(BarePath, extractZipName), filepath.Join(BarePath, parser.BareObj.BareName))
	if err != nil {
		log.Fatal("Error renaming template")
	}
}
