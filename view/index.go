package view

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"path/filepath"

	"github.com/chunkhang/twocents/util"
	"github.com/labstack/echo"
)

const (
	layoutDirectory = "template/layout"
	pageDirectory   = "template/page"
	fileExtension   = ".tmpl"
)

var (
	mainTemplate = template.Must(template.New("main").
		Parse(`{{define "main" }} {{template "base" .}} {{end}}`))
)

// Registry is the registry of templates
type Registry struct {
	templates map[string]*template.Template
}

// Render renders the page with given data
func (t *Registry) Render(w io.Writer, name string, data interface{}, c echo.Context) (err error) {
	template, ok := t.templates[name]
	if ok {
		err = template.ExecuteTemplate(w, "main", data)
	} else {
		err = errors.New("Template not found: " + name)
	}
	return
}

// Init builds all the templates
func Init() (r *Registry, err error) {
	defer util.Catch(&err)
	r = &Registry{}

	layoutFiles, err := getFiles(layoutDirectory, fileExtension)
	util.Check(err)

	pageFiles, err := getFiles(pageDirectory, fileExtension)
	util.Check(err)

	templates := make(map[string]*template.Template)
	for _, file := range pageFiles {
		filename := filepath.Base(file)
		files := append(layoutFiles, file)
		m := template.Must(mainTemplate.Clone())
		templates[filename] = template.Must(m.ParseFiles(files...))
	}
	r.templates = templates

	return
}

func getFiles(directory string, extension string) (files []string, err error) {
	pattern := fmt.Sprintf("%s/*%s", directory, extension)
	files, err = filepath.Glob(pattern)
	return
}
