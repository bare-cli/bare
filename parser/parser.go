package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// func Parser(filePath string) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		fmt.Println("File not found")
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		inputCmd := strings.TrimSpace(scanner.Text())
// 		inputCmdArg := strings.Split(inputCmd, " ")
// 		RecipeCmds(inputCmdArg)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 	}
// }

type Bare struct {
	BareName string
	Version string
	Include []string
	BarePath string 
}

var BareObj Bare

func Parser(filePath string){
	data, err := ioutil.ReadFile(filePath);
	if err != nil {
		log.Fatal(err);
	}

	// var bareObj Bare;
	err = json.Unmarshal(data, &BareObj)
	if err != nil {
		log.Fatal(err);
	}

	if BareObj.BareName == "" {
		log.Fatal("Bare name not present")
	}
}