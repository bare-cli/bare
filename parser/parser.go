package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parser(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputCmd := strings.TrimSpace(scanner.Text())
		inputCmdArg := strings.Split(inputCmd, " ")
		RecipeCmds(inputCmdArg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

// func Parser(filePath string){
// 	data, err := ioutil.ReadFile(filePath);
// 	if err != nil {
// 		log.Fatal(err);
// 	}


// }