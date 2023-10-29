package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Req struct {
	*http.Request
}

func wrapReq(r *http.Request) *Req {
	return &Req{r}
}

func (r *Req) urlParam(param string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return chi.URLParam(r.Request, param), nil
}

func (r *Req) urlParamInt(param string) (int, error) {
	p, err := r.urlParam(param)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(p)
}

type Res struct {
	http.ResponseWriter
}

func wrapRes(w http.ResponseWriter) *Res {
	return &Res{w}
}

func (r *Res) pushUrl(path string) {
	r.Header().Add("HX-Push-Url", path)
}
