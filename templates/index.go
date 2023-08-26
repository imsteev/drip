package templates

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *.tmpl
var templateFS embed.FS

var tmpl *template.Template

type Index struct {
	Messages []string
}

func (it Index) Render(w io.Writer) error {
	if tmpl == nil {
		var err error
		if tmpl, err = template.New("index").
			ParseFS(templateFS, "*.tmpl"); err != nil {
			return err
		}
	}
	return tmpl.ExecuteTemplate(w, "base.tmpl", it.Messages)
}
