package web

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/app"
	wiki "github.com/mmm888/go-wiki/domain"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type EditGetHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
	Service    *app.EditService
}

func (h *EditGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in := &wiki.EditInput{
		CommonVars: h.CommonVars,
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
	tmpl := template.Must(template.New("edit.tmpl").Funcs(funcMap).ParseFiles("templates/edit.tmpl"))
	if err := tmpl.Execute(w, out); err != nil {
		log.Fatal(err)
	}
}

type EditPostHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
	Service    *app.EditService
}

func (h *EditPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//var err error

	in := &wiki.EditInput{
		CommonVars: h.CommonVars,
	}

	q := r.URL.Query()

	if v := q.Get("path"); v != "" {
		in.Path = v
	}

	if v := r.FormValue("name"); v != "" {
		in.Name = v
	}

	if v := r.FormValue("type"); v != "" {
		in.FileType = v
	}

	if v := r.FormValue("content"); v != "" {
		in.Contents = v
	}

	if err := h.Service.Info.Post(in); err != nil {
		log.Print(err)
	}

	redirect := fmt.Sprintf("/show?path=%s", in.Path)
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}
