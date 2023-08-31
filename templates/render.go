package templates

import (
	"drip/data"
	"embed"
	"fmt"
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
	Messages []*data.Message
	RoomURL  string
	Space    string

	DeleteURL string
}

func (it Index) Render(w io.Writer) error {
	it.DeleteURL = fmt.Sprintf("/spaces/%s/drip", it.Space)
	return tmpl.ExecuteTemplate(w, "base.tmpl", it)
}
