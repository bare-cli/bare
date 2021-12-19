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
	"strings"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var keepDownloadedZip *bool = flag.Bool("keep", false, "Keep downloaded ")

func init() {
	rootCmd.AddCommand(useCmd)
	flag.Parse()
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
	Name         string
	Template     string
	Placeholders map[string]string
	BarePath     string
}

var TempObj NewTemplate

func useBare(bareName, desti string) {
	user, repo, branch := parser.ParseGithubRepo(bareName)
	parser.GetRecipe(user, repo, branch)

	// Prompt project name and template
	TempObj.Template = ui.VarPromptSelect("Select template", parser.BareObj.Variants)

	// Prompt placeholders
	TempObj.Placeholders = make(map[string]string)
	for k, e := range parser.BareObj.Placeholders {
		TempObj.Placeholders[k] = ui.PromptString(k, e)
	}
	osutil.MakeDownloadFolder()
	err := git.DownloadZip(user, repo, branch, parser.BareObj.BareName)
	if err != nil {
		log.Fatal(err)
	}

	downloadZipName := parser.BareObj.BareName + ".zip"
	downloadZipPath := filepath.Join(BarePath, "tmp", downloadZipName)
	extractZipName := parser.BareObj.BareName + "_" + user
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
	err = template.Execute(varPath, targetPath, TempObj.Placeholders)
	if err != nil {
		log.Fatal(err)
	}

	if !*keepDownloadedZip {
		if osutil.Exists(filepath.Join(extractZipPath, extractZipName)) {
			err = os.RemoveAll(filepath.Join(extractZipPath, extractZipName))
		}
	}

	// Git init
	gitDefVal := []string{"n", "y/N"}
	var gitInitData git.NewRepo
	doGitInitPrompt := ui.PromptString("Initialize with git", gitDefVal)

	if strings.ToLower(doGitInitPrompt) != "n" {
		gitInitData.Path = filepath.Join(targetPath, TempObj.Placeholders["AppName"])
		gitInitData.Remote = ui.PromptString("Git remote link", []string{"", "Enter link to online repository"})
		if gitInitData.Remote == "" {
			fmt.Println(styles.Error.Render("Remote link cannot be empty"))
			os.Exit(1)
		}
		err = git.GitInit(gitInitData)
		if err != nil {
			fmt.Println(styles.Error.Render("Error with git init!"))
			os.Exit(1)
		}
	}

	// Add license
	git.AddLicense(filepath.Join(targetPath, TempObj.Placeholders["AppName"], "LICENSE"))

	fmt.Println("Your project has been created", styles.InitStyle.Render("GLHF!!"))
}
