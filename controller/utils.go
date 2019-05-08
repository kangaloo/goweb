package controller

import (
	"html/template"
	"io/ioutil"
	"os"
)

func populateTemplates() map[string]*template.Template {
	basePath := "templates"
	res := make(map[string]*template.Template)
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {
		func() {
			f, err := os.Open(basePath + "/content/" + fi.Name())
			if err != nil {
				panic("Failed to open template '" + err.Error() + "'")
			}
			defer func() { _ = f.Close() }()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				panic("Failed to read content from file '" + fi.Name() + "'")
			}

			tpl := template.Must(layout.Clone())
			_, err = tpl.Parse(string(content))
			if err != nil {
				panic("Failed to parse contents of '" + fi.Name() + "' as template")
			}

			res[fi.Name()] = tpl
		}()
	}

	return res
}
