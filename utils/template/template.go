package template

import (
	"os"
	"path/filepath"
	"text/template"

	"bare/utils/stringutils"
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

// Replaces {{AppName}} from the template
func Execute(path string, fileName string) error {
	return filepath.Walk(path, func(fileName string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		oldName, err := filepath.Rel(path, fileName)
		if err != nil {
			return err
		}

		buf := stringutils.NewString("")
		appName := AppNameReplace{"working"}
		fnameTmpl := template.Must(template.New("File name template").Option(Options...).Parse(oldName))
		if err := fnameTmpl.Execute(buf, appName); err != nil {
			return err
		}
		return nil
	})
}
