package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd);
}
var addCmd = &cobra.Command{
	Use: "add",
	Short: "add you current bare to repo",
	Long: "same as short for addCmd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Reading your recipe ...")
	},
}
