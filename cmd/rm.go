package cmd

import (
	"bare/styles"
	"bare/utils/osutil"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "delete an exsisting bare",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1 {
			rmBare(args)
		} else {
			fmt.Println(styles.Error.Render("Not enought arguments"))
		}
	},
}

func rmBare(delBares []string) {
	fmt.Println(styles.InitStyle.Render("~ bare rm ~"))
	barePath := filepath.Join(os.Getenv("HOME"), ".bare")
	deletedBare := []string{}
	for _, bare := range delBares {
		if osutil.Exists(filepath.Join(barePath, bare)) {
			fmt.Println(styles.StatusError.Render("X"), bare)
			err := os.RemoveAll(filepath.Join(barePath, bare))
			if err != nil {
				fmt.Println(styles.Error.Render("[Error] There was problem removing"), bare)
				os.Exit(1)
			}
			deletedBare = append(deletedBare, bare)
		} else {
			fmt.Println(styles.Warning.Render("Bare does not exsist!"), bare)
		}
	}
	if len(deletedBare) > 0 {
		fmt.Println(styles.Success.Render("[Success] successfully removed"), strings.Trim(fmt.Sprint(deletedBare), "[]"))
	}

}
