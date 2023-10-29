package main

import (
	"fmt"
	"net/http"
)

type Res struct {
	http.ResponseWriter
}

func wrapRes(w http.ResponseWriter) *Res {
	return &Res{w}
}

func (r *Res) pushUrl(path string) {
	r.Header().Add("HX-Push-Url", path)
}

func (r *Res) writef(format string, params ...any) {
	r.Write([]byte(fmt.Sprintf(format, params...)))
}
