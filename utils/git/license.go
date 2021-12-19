package git

import (
	"bare/utils/ui"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type NewLicense struct {
	LicName string
	Path    string
	LicStr  string
}

type LicenseTempl struct {
	License string
}

var License NewLicense

func AddLicense(path string) {
	doAddLicense := ui.PromptString("Add license", []string{"n", "y/N"})
	if strings.ToLower(doAddLicense) == "n" {
		return
	}
	licName := []string{"Apache Software License 2.0", "MIT", "GNU GPL v3.0"}
	License.LicName = ui.VarPromptSelect("Select license", licName)
	License.LicStr = getLicense()
	License.Path = path
	executeLicense()
}

func getLicense() string {
	req_url := "https://raw.githubusercontent.com/bare-cli/license/main/LICENSE"
	resp, err := http.Get(req_url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func executeLicense() {

	f, err := os.OpenFile(License.Path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fileTemplateName := filepath.Base(License.Path)
	licTempl := template.Must(template.New(fileTemplateName).Parse(License.LicStr))

	licData := LicenseTempl{
		License: License.LicName,
	}
	err = licTempl.ExecuteTemplate(f, fileTemplateName, licData)
	if err != nil {
		log.Fatal(err)
	}
}
