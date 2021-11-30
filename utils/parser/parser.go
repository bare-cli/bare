package parser

import (
	"bare/utils/osutil"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const raw_url = "https://raw.githubusercontent.com/"

type Bare struct {
	BareName     string
	Version      string
	BarePath     string
	Variants     []string // Template name -> description (to be asked in prompt)
	Placeholders map[string]string
}

var BareObj Bare

func Parser(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// var bareObj Bare;
	err = json.Unmarshal(data, &BareObj)
	if err != nil {
		log.Fatal(err)
	}

	if BareObj.BareName == "" {
		log.Fatal("Bare name not present")
	}
}

func UpdateRecipe() {
	currDir, _ := os.Getwd()
	recipePath := currDir + "/recipe.json"
	updatedRecipe, err := json.MarshalIndent(BareObj, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	if osutil.Exists(recipePath) {
		err = ioutil.WriteFile(recipePath, updatedRecipe, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetRecipe(user string, repo string, branch string) {
	req_url := raw_url + user + "/" + repo + "/" + branch + "/recipe.json"
	resp, err := http.Get(req_url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	err = json.Unmarshal(body, &BareObj)
	if err != nil {
		log.Fatal(err)
	}
}

// Return user, repo, branch
func ParseGithubRepo(bareName string) (string, string, string) {

	userRepo := strings.Split(bareName, "/")

	if len(userRepo) <= 1 {
		fmt.Println("Invalid Bare path")
		os.Exit(-1)
	}

	user := userRepo[0]
	repo := userRepo[1]

	repoBranch := strings.Split(repo, "#")
	branch := "main"
	if len(repoBranch) > 1 {
		branch = repoBranch[1]
	}

	repo = repoBranch[0]

	return user, repo, branch
}
