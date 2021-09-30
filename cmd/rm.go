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
	rootCmd.AddCommand(rmCmd)
}

var rmCmd = &cobra.Command{
	Use: "rm",
	Short: "delete an exsisting bare" ,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 1{
			rmBare(args)
		}else{
			fmt.Println(styles.InitError.Render("Not enought arguments"))
		}
	},
}

func rmBare(delBares []string ){
	fmt.Println(styles.InitStyle.Render("Bare rm"))
	barePath := filepath.Join(os.Getenv("HOME"), ".bare")
	for _, bare := range delBares {
		if utils.Exists(filepath.Join(barePath, bare)) {
			fmt.Println(styles.InitError.Render("[Deleting] "), bare)
			os.RemoveAll(filepath.Join(barePath, bare))
		}
	} 
}
