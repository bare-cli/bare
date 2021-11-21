package cmd

import (
	"bare/styles"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var BarePath = filepath.Join(os.Getenv("HOME"), ".bare")

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
