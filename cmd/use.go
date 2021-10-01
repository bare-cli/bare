package cmd

import (
	"bare/styles"
	"bare/utils"
	"fmt"
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

func useBare(bareName, desti string) {
	currDir, _ := os.Getwd()
	barePath := filepath.Join(os.Getenv("HOME"), ".bare")

	if !utils.Exists(filepath.Join(barePath, bareName)) {
		fmt.Println(styles.InitError.Render("Bare doesn't exsist"))
		fmt.Println("User `bare list` to get list of all the bares")
		os.Exit(1)
	}

	if utils.Exists(filepath.Join(currDir, desti)) {
		fmt.Println(styles.InitError.Render("File name already exsists"))
		os.Exit(1)
	} else {
		utils.CreateIfNotExists(filepath.Join(currDir, desti), 0755)
		utils.CopyDirectory(filepath.Join(barePath, bareName), filepath.Join(currDir, desti))
	}

}
