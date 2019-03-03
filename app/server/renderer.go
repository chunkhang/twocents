package server

import (
	"errors"
	"html/template"
	"io"
	"path/filepath"

	"github.com/chunkhang/twocents/app/util"
	"github.com/labstack/echo"
)

const (
	layoutDirectory = "view/layout"
	pageDirectory   = "view/page"
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
	defer util.Catch(&err)

	template, ok := t.templates[name]
	if ok {
		err = template.ExecuteTemplate(w, "main", data)
	} else {
		err = errors.New("Template not found: " + name)
	}
	util.Check(err)

	return
}

func (s *server) setRenderer() (err error) {
	defer util.Catch(&err)
	r := &Registry{}

	layoutFiles, err := util.GetFiles(layoutDirectory, fileExtension)
	util.Check(err)

	pageFiles, err := util.GetFiles(pageDirectory, fileExtension)
	util.Check(err)

	templates := make(map[string]*template.Template)
	for _, file := range pageFiles {
		filename := filepath.Base(file)
		files := append(layoutFiles, file)
		m := template.Must(mainTemplate.Clone())
		templates[filename] = template.Must(m.ParseFiles(files...))
	}
	r.templates = templates

	s.Renderer = r

	return
}
