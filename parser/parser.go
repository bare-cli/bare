package parser

import (
	"bare/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Bare struct {
	BareName string
	Version string
	Include []string
	BarePath string
	Touch map[string]string
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

func UpdateRecipe() {
	currDir, _ := os.Getwd()
	recipePath := currDir + "/recipe.json"
	updatedRecipe, err := json.MarshalIndent(BareObj, "", "    ");
	if err != nil {
		log.Fatal(err)
	}
	if utils.Exists(recipePath) {
		err = ioutil.WriteFile(recipePath, updatedRecipe, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}