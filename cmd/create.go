package cmd

import "github.com/spf13/cobra"

func init(){
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use: "create",
	Short: "Creates an empty template for your bare..",
	Run: func(cmd *cobra.Command, args []string) {
		createTemplate();
	},
}

func createTemplate() {
	// TODO
	// [ ] Get template from net and edit
	// [ ] Generate an empty template 
}