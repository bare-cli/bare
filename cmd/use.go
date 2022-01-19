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

var keepDownloadedZip *bool = flag.Bool("keep", false, "Keep downloaded.")
var shouldUseDefault *bool = flag.Bool("d", false, "Use default value while creating project.")
var shouldInitGit *bool = flag.Bool("git", false, "Initialize with git version control.")

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
	TempObj.Placeholders = make(map[string]string)

	// Set placeholders, prompt if the default flag is not used.
	if !*shouldUseDefault {
		for k, e := range parser.BareObj.Placeholders {
			TempObj.Placeholders[k] = ui.PromptString(k, e)
		}
	} else if *shouldUseDefault {
		for k, e := range parser.BareObj.Placeholders {
			TempObj.Placeholders[k] = e[0]
		}
	}
	fmt.Println(TempObj)
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
	currDir, _ := os.Getwd()
	targetPath := filepath.Join(currDir, desti)
	err = template.Execute(varPath, targetPath, TempObj.Placeholders)
	if err != nil {
		log.Fatal(err)
	}

	if !*keepDownloadedZip {
		if osutil.Exists(filepath.Join(extractZipPath, extractZipName)) {
			err = os.RemoveAll(filepath.Join(extractZipPath, extractZipName))
			if err != nil {
				fmt.Println(styles.Error.Render("Error deleting the downloaded codebase."))
				os.Exit(1)
			}
		}
	}

	// Git init
	gitDefVal := []string{"n", "y/N"}
	var gitInitData git.NewRepo
	var doGitInitPrompt = "y"
	if !*shouldInitGit {
		doGitInitPrompt = ui.PromptString("Initialize with git", gitDefVal)
	}

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
	if !*shouldUseDefault {
		git.AddLicense(filepath.Join(targetPath, TempObj.Placeholders["AppName"], "LICENSE"))
	}

	fmt.Println("Your project has been created", styles.InitStyle.Render("GLHF!!"))
}
