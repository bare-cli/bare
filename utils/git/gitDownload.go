package git

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const zipDownloadUrl = "https://codeload.github.com/"
const zipDownloadFormat = "/zip/"

func DownloadZip(user, repo, branch, fileName string) error {
	zipUrl := zipDownloadUrl + user + "/" + repo + zipDownloadFormat + branch

	resp, err := http.Get(zipUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Status Code is not 200")
	}
	homePath := os.Getenv("HOME")
	out, err := os.Create(filepath.Join(homePath, ".bare", "tmp", fileName+".zip"))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
