package template

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Registry is the registry of templates
type Registry struct {
	templates map[string]*template.Template
}

// Render attempts to render the template
func (t *Registry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

// Init initializes the templates
func Init() *Registry {
	t := &Registry{}

	templates := make(map[string]*template.Template)
	templates["home.html"] = template.Must(template.ParseFiles("view/home.tmpl", "view/base.tmpl"))

	t.templates = templates

	return t
}
