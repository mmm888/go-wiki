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

type ConfigGetHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
	Service    *app.ConfigService
}

func (h *ConfigGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in := &wiki.ConfigInput{
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

	if out.Path != "" {
		out.Query = fmt.Sprintf("?path=%s", out.Path)
	}

	funcMap := template.FuncMap{}
	tmpl := template.Must(template.New("config.tmpl").Funcs(funcMap).ParseFiles("templates/config.tmpl"))
	if err := tmpl.Execute(w, out); err != nil {
		log.Fatal(err)
	}
}

type ConfigPostHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
	Service    *app.ConfigService
}

func (h *ConfigPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	in := &wiki.ConfigInput{
		CommonVars: h.CommonVars,
	}

	if v := r.FormValue("name"); v != "" {
		in.Name = v
	}

	if v := r.FormValue("repo"); v != "" {
		in.Repo = v
	}

	if err := h.Service.Info.Post(in); err != nil {
		log.Print(err)
	}

	http.Redirect(w, r, "/config", http.StatusSeeOther)
}
