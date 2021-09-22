package cmd

import (
	"bare/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd);
}

var initCmd = &cobra.Command{
	Use: "init",
	Short: "inits the project with bare",
	Run: func(cmd *cobra.Command, args []string) {
		bareInit(args[0]);
	},
} 


func bareInit(bareName string){

	currDir, _ := os.Getwd();
	homePath := os.Getenv("HOME");

	bareFolderExsists := utils.CheckFolder(homePath + "/.bare/" + bareName);
	recipeFileExsists := utils.CheckFolder(currDir + "/Recipe");

	if bareFolderExsists {
		log.Fatal("Bare with similar name exsists")
	}else if !bareFolderExsists{
		if !recipeFileExsists {
			if _, err := os.Create(currDir + "/Recipe"); err != nil {
				log.Fatal(err);
			}
		}
	
		if err := os.Mkdir(homePath + "/.bare/" + bareName, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}