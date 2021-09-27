package cmd

import (
	"bare/parser"
	"bare/styles"
	"bare/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "inits the project with bare",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.InitStyle.Render("Bare Init"))
		var bareName string
		if len(args) == 0 {
			fmt.Println(styles.InitError.Render("[Error] Provide bare name > bare init <bare-name>"))
			os.Exit(1)
		} else {
			bareName = args[0]
		}
		bareInit(bareName)
	},
}

func bareInit(bareName string) {

	currDir, _ := os.Getwd()
	homePath := os.Getenv("HOME")

	bareFolderExsists := utils.Exists(homePath + "/.bare/" + bareName)
	recipeFileExsists := utils.Exists(currDir + "/recipe.json")

	if bareFolderExsists {
		fmt.Println(styles.InitError.Render("[Error] Bare with similar name exsists"))
		os.Exit(1)
	} else if !bareFolderExsists {
		if !recipeFileExsists {
			newBare := parser.Bare{
				BareName: bareName,
				BarePath: homePath + "/.bare/" + bareName,
			}
			res, err := json.MarshalIndent(newBare, "", "    ")
			if err != nil {
				log.Fatal(err)
			}
			recipeErr := ioutil.WriteFile(currDir+"/recipe.json", res, os.ModePerm)
			if recipeErr != nil {
				fmt.Println(styles.InitError.Render("[Error] Cannot create recipe.json file"))
			} else {
				fmt.Println(styles.InitSuccess.Render("[Success] Created recipe.json"))
			}
		}

		if err := os.Mkdir(homePath+"/.bare/"+bareName, os.ModePerm); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(styles.InitSuccess.Render("[Success] Created new bare "), bareName)
		}
	}
}
