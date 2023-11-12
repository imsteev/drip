package templates

import (
	"drip/data/models"
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
	Messages []*models.Message
	RoomURL  string
	SpaceID  int
}

func (it Index) Render(w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "entrypoint.tmpl", it)
}

type Share struct {
	RoomURL string
	SpaceID int
}

func (s Share) Render(w io.Writer) error {
	return tmpl.ExecuteTemplate(w, "share.tmpl", s)
}
