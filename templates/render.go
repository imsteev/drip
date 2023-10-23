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

func (it Index) MustRender(w io.Writer) {
	_ = it.Render(w)
}
