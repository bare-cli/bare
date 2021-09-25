package cmd

import (
	"bare/styles"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use : "list",
	Short : "List all the saved bare",
	Run : func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.InitStyle.Render("Lists of all the bare"))
		if len(args) == 0{
			ListBare()
		}else {
			ListFileWalk(args[0])
		}
	},
}


func ListBare() {
	barePath := os.Getenv("HOME") + "/.bare"
	

	fis, err := ioutil.ReadDir(barePath)

	if err != nil {
		fmt.Println(styles.InitError.Render("Error reading directory from ~/.bare"))
		os.Exit(1)
	}

	for _, info := range fis {
		if info.Name()[0] != '$' {
			fmt.Printf("|_%v\n", info.Name())
		}
	}
}

func ListFileWalk(bareName string) {
	barePath := os.Getenv("HOME") + "/.bare/" + bareName
	err := filepath.Walk(barePath,  func(path string, info fs.FileInfo, err error) error {
		if err != nil{
			return err
		}
		fmt.Println(path)
		return nil 
	})
	if err != nil {
		log.Fatal(err)
	}
}