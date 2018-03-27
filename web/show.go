package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/app"
	wiki "github.com/mmm888/go-wiki/domain"
	"github.com/mmm888/go-wiki/middleware/markdown"
	"github.com/mmm888/go-wiki/middleware/templates"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type ShowHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
	Templates  *templates.Templates
	Markdown   markdown.Markdown
	Service    *app.ShowService
}

func (h *ShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//var err error

	in := &wiki.ShowInput{
		CommonVars: h.CommonVars,
		Markdown:   h.Markdown,
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

	if err := h.Templates.Render("show", w, out); err != nil {
		log.Print(err)
	}
}
