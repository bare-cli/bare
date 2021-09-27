package cmd

import (
	"bare/styles"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the saved bare",
	Run: func(cmd *cobra.Command, args []string) {
		
		if len(args) == 0 {
			ListBare()
		} else {
			ListFileWalk(args[0])
		}
	},
}

func ListBare() {
	fmt.Println(styles.InitStyle.Render("Lists of all the bare"))
	barePath := os.Getenv("HOME") + "/.bare"

	fis, err := ioutil.ReadDir(barePath)

	if err != nil {
		fmt.Println(styles.InitError.Render("Error reading directory from ~/.bare"))
		os.Exit(1)
	}

	for ind , info := range fis {
		if info.Name()[0] != '$' {
			if ind != len(fis) - 1 {
				fmt.Printf("├──%v\n", info.Name())
			}else {
				fmt.Printf("└──%v\n", info.Name())
			}
		}
	}
}

func ListFileWalk(bareName string) {
	fmt.Println(styles.InitStyle.Render("Lists of all the files in"), styles.InitSuccess.Render(bareName))
	barePath := os.Getenv("HOME") + "/.bare/" + bareName
	printDirectory(barePath, 0)
}

func printListing(entry string, depth int) {
	indent := strings.Repeat("│   ", depth)
	fmt.Printf("%s├── %s\n", indent, entry)
}

func printDirectory(path string, depth int) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	printListing(path, depth)
	for _, entry := range entries {
		if (entry.Mode() & os.ModeSymlink) == os.ModeSymlink {
			full_path, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Printf("error reading link: %s\n", err.Error())
			} else {
				printListing(entry.Name()+" -> "+full_path, depth+1)
			}
		} else if entry.IsDir() {
			printDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {
			printListing(entry.Name(), depth+1)
		}
	}
}