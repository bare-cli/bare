package git

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const zipDownloadUrl = "https://codeload.github.com/"

func DownloadZip(user, repo, branch, fileName string) error {
	zipUrl := zipDownloadUrl + user + "/" + repo + "/" + "tar.gz/" + branch

	resp, err := http.Get(zipUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Status Code is not 200")
	}
	homePath := os.Getenv("HOME")
	fmt.Println("#####", fileName)
	out, err := os.Create(filepath.Join(homePath, ".bare", "tmp", fileName+".tar.gz"))
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
