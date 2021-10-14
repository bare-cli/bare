package cmd

import (
	"bare/styles"
	"bare/utils"
	"bare/utils/host"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "gets a project template from a github repository to template repository",
	Run: func(cmd *cobra.Command, args []string) {
		getGithub(args)
	},
}

func getGithub(args []string ) {
	// TODO : validate template name and url
	// args[0] => template url
	// args[1] => destination name
	currDir, _ := os.Getwd()
	targetPath := path.Join(currDir, args[1])

	if utils.Exists(path.Join(currDir, args[1])) {
		fmt.Println(styles.InitError.Render("Directory with similar name already exsists !"))
		os.Exit(1)
	}

	if err := utils.CloneRepo(targetPath, utils.CloneOptions{
		URL : host.URL(args[0]),
	}); err != nil {
		fmt.Println(styles.InitError.Render("There was error cloning the repo"), args[0])
	}
}