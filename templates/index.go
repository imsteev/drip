package templates

import (
	"html/template"
	"io"
)

var (
	tmpl *template.Template
)

type IndexTemplate struct {
	Messages []string
}

func (it IndexTemplate) Render(w io.Writer) error {
	if tmpl == nil {
		var err error
		if tmpl, err = template.ParseGlob("./templates/*.tmpl"); err != nil {
			return err
		}
	}
	return tmpl.ExecuteTemplate(w, "base.tmpl", it.Messages)
}
