package osutil

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func UnzipReader(r io.Reader, source, target string) {
	reader := bufio.NewReader(r)
	src := new(bytes.Buffer)
	_, err := src.ReadFrom(reader)
	if err != nil {
		log.Fatal(err)
	}
	of, err := os.OpenFile(source, os.O_CREATE|os.O_WRONLY, 0777)
	defer of.Close()

	if err != nil {
		log.Fatal(err)
	}
	dst := new(bytes.Buffer)
	src.ReadFrom(reader)
	_, err = io.Copy(dst, src)

	if err != nil {
		log.Fatal(err)
	}

	_, err = dst.WriteTo(of)
	if err != nil {
		log.Fatal(err)
	}

	Unzip(source, target)
}

func Unzip(fileName, target string) {
	r, err := zip.OpenReader(fileName)
	defer r.Close()
	if err != nil {
		log.Fatal(err)
	}

	os.MkdirAll(target, 0777)
	for _, it := range r.File {
		name := it.Name
		n := strings.Index(name, "/")
		dstName := name[n+1:]
		if len(dstName) <= 0 {
			continue
		}

		dstName = filepath.Join(target, dstName)
		if it.FileInfo().IsDir() {
			err := os.MkdirAll(dstName, 0777)
			if err != nil {
				continue
			}
			continue
		}

		dstF, err := os.Create(dstName)
		defer dstF.Close()

		if err != nil {
			continue
		}

		src := new(bytes.Buffer)
		dst := new(bytes.Buffer)
		rc, err := it.Open()
		defer rc.Close()
		if err != nil {
			continue
		}
		src.ReadFrom(rc)
		_, err = io.Copy(dst, src)
		if err != nil {
			continue
		}

		dst.WriteTo(dstF)

	}
}
