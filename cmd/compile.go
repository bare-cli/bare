package cmd

import (
	"bare/parser"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(compileCommand)
}

var compileCommand = &cobra.Command{
	Use: "compile",
	Run: func(cmd *cobra.Command, args []string) {
		parser.Parser(args[0])
	},
}
