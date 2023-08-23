package templates

import (
	"html/template"
	"io"
)

var (
	indexHtmlTmpl *template.Template
)

type IndexTemplate struct {
	Messages []string
}

func (it IndexTemplate) Render(w io.Writer) error {
	if indexHtmlTmpl == nil {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			return err
		}
		indexHtmlTmpl = tmpl
	}
	return indexHtmlTmpl.Execute(w, it.Messages)
}

type MeTemplate struct {
	Messages []string
}

func (mt MeTemplate) Render(w io.Writer) error {
	tmpl, err := template.ParseFiles("./templates/me.html")
	if err != nil {
		return err
	}
	return tmpl.Execute(w, mt.Messages)
}
