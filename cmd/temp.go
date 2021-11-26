package cmd

import (
	"bare/utils/template"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(templateCmd)
}

var templateCmd = &cobra.Command{
	Use:   "temp",
	Short: "add you current bare to repo",
	Long:  "same as short for addCmd",
	Run: func(cmd *cobra.Command, args []string) {
		tempBare()
	},
}

func tempBare() {
	fmt.Println("Hello world")
	bareName := "test"
	// currDir, _ := os.Getwd()
	barePath := filepath.Join(os.Getenv("HOME"), ".bare", bareName)
	err := template.Execute(barePath)
	if err != nil {
		log.Fatal(err)
	}
}
