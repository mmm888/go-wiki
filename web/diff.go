package web

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/app"
	wiki "github.com/mmm888/go-wiki/domain"
	"github.com/mmm888/go-wiki/middleware/git"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type DiffHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
	Service    *app.DiffService
}

func (h *DiffHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in := &wiki.DiffInput{
		CommonVars: h.CommonVars,
		Git:        git.NewGit(h.CommonVars.Name, h.CommonVars.Repo),
	}

	q := r.URL.Query()

	if v := q.Get("path"); v != "" {
		in.Path = v
	}

	out, err := h.Service.Info.Get(in)
	if err != nil {
		log.Print(err)
	}

	if out == nil {
		return
	}

	if out.Path != "" {
		out.Query = fmt.Sprintf("?path=%s", out.Path)
	}

	funcMap := template.FuncMap{}
	tmpl := template.Must(template.New("diff.tmpl").Funcs(funcMap).ParseFiles("templates/diff.tmpl"))
	if err := tmpl.Execute(w, out); err != nil {
		log.Print(err)
	}
}
