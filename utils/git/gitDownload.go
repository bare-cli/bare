package git

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const zipDownloadUrl = "https://codeload.github.com/"

func DownloadZip(user, repo, branch, fileName string) {
	zipUrl := zipDownloadUrl + user + "/" + repo + "/" + "tar.gz/" + branch

	resp, err := http.Get(zipUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return
	}
	homePath := os.Getenv("HOME")
	out, err := os.Create(filepath.Join(homePath, ".bare", "tmp", fileName+".tar.gz"))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
