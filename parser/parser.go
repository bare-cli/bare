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
		inputCmdArg := strings.Split(inputCmd, " ");
		fmt.Println(inputCmdArg[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}