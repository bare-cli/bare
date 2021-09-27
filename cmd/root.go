package cmd

import (
	"bare/styles"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bare",
	Short: "manager for you barebones",
	Long:  "manage for your barebones",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.InitStyle.Render("Welcome to Bare, use -h flag for all the commands"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
