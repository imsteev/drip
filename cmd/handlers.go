package main

import (
	"drip/data"
	"drip/templates"
	"drip/utils"
	"net/http"
)

func GetMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := templates.Index{
		Messages: s.GetMessages(data.MY_SPACE),
	}
	if err := tmpl.Render(w); err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}
}

func CreateDrip(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteStrf(w, "form error: %v", err)
		return
	}
	s.AddMessage(r.FormValue("text"), data.MY_SPACE)

	tmpl := templates.Index{
		Messages: s.GetMessages(data.MY_SPACE),
	}
	if err := tmpl.Render(w); err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}
}

func DeleteDrip(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteStrf(w, "form error: %v", err)
		return
	}
	text := r.FormValue("text")
	s.DeleteMessage(text, data.MY_SPACE)

}
