package cmd

import (
	"bare/styles"
	"bare/utils/git"
	"bare/utils/host"
	"bare/utils/osutil"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets a project template from a github repository to template repository",
	Run: func(cmd *cobra.Command, args []string) {
		getGithub(args)
	},
}

func getGithub(args []string) {
	// TODO : validate template name and url
	// args[0] => template url
	// args[1] => destination name
	currDir, _ := os.Getwd()
	targetPath := path.Join(currDir, args[1])

	if osutil.Exists(path.Join(currDir, args[1])) {
		fmt.Println(styles.Error.Render("Directory with similar name already exsists !"))
		os.Exit(1)
	}

	if err := git.CloneRepo(targetPath, git.CloneOptions{
		URL: host.URL(args[0]),
	}); err != nil {
		fmt.Println(styles.Error.Render("There was error cloning the repo"), args[0])
	}
}
