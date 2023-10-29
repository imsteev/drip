package main

import (
	"io"
	"net/http"
)

type Res struct {
	http.ResponseWriter
}

type Renderer interface {
	Render(w io.Writer) error
}

func wrapRes(w http.ResponseWriter) *Res {
	return &Res{w}
}

func (r *Res) pushUrl(path string) {
	r.Header().Add("HX-Push-Url", path)
}

func (r *Res) render(rr Renderer) error {
	return rr.Render(r)
}
