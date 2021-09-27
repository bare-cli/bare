package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
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