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

func init() {
	tmpl = template.Must(template.New("templates").ParseFS(templateFS, "*.tmpl"))
}

type Index struct {
	Messages []string
	RoomURL  string
	Space    string
}

func (it Index) Render(w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "base.tmpl", it)
}
