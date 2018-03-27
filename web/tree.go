package web

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/app"
	wiki "github.com/mmm888/go-wiki/domain"
	"github.com/mmm888/go-wiki/middleware/templates"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type TreeHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
	Templates  *templates.Templates
	Service    *app.TreeService
}

func (h *TreeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//var err error

	in := &wiki.TreeInput{
		CommonVars: h.CommonVars,
	}

	out, err := h.Service.Info.Get(in)
	if err != nil {
		log.Print(err)
	}

	if out == nil {
		return
	}

	if err := h.Templates.Render("tree", w, out); err != nil {
		log.Print(err)
	}
}
