package template

import (
	"bare/utils/stringutils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"text/template"
)

type Interface interface {
	Execute(string) error

	UseDefaultValues()
}

type Template struct {
	Path              string
	ShouldUseDefaults bool
}

/*
- download the template
- run through the template and change the {{AppName}}
*/

var Options = []string{
	"missingkey=invalid",
}

type AppNameReplace struct {
	AppName string
}

// Replaces all the placeholders from the template
func Execute(source string, dirPrefix string, placeholders map[string]string) error {

	isOnlyWhitespace := func(buf []byte) bool {
		wsre := regexp.MustCompile(`\S`)

		return !wsre.Match(buf)
	}

	return filepath.Walk(source, func(fileName string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		oldName, err := filepath.Rel(source, fileName)
		if err != nil {
			return err
		}

		buf := stringutils.NewString("")
		fnameTmpl := template.Must(template.New("File name template").Option(Options...).Parse(oldName))
		if err := fnameTmpl.Execute(buf, placeholders); err != nil {
			return err
		}
		newName := buf.String()

		target := filepath.Join(dirPrefix, newName)

		if info.IsDir() {
			if err := os.Mkdir(target, 0755); err != nil {
				if !os.IsExist(err) {
					return err
				}
			}
		} else {
			fi, err := os.Lstat(fileName)
			if err != nil {
				return err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, fi.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			defer func(fname string) {
				contents, err := ioutil.ReadFile(fname)
				if err != nil {
					fmt.Printf("couldn't read the contents of file %q, got error %q", fname, err)
					return
				}

				if isOnlyWhitespace(contents) {
					os.Remove(fname)
					return
				}
			}(f.Name())
			contentsTmpl := template.Must(template.New("File contents template").Option(Options...).ParseFiles(fileName))
			fileTemplateName := filepath.Base(fileName)

			if err := contentsTmpl.ExecuteTemplate(f, fileTemplateName, placeholders); err != nil {
				return err
			}
		}
		return nil
	})
}
