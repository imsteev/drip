package templates

import (
	"embed"
	"html/template"
	"io"
)

var (
	tmpl *template.Template

	//go:embed *.tmpl
	templateFS embed.FS
)

type Index struct {
	Messages []string
	RoomURL  string
}

func (it Index) Render(w io.Writer) error {
	if tmpl == nil {
		var err error
		if tmpl, err = template.New("index").
			ParseFS(templateFS, "*.tmpl"); err != nil {
			return err
		}
	}
	return tmpl.ExecuteTemplate(w, "base.tmpl", it)
}
