package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type RootHandler struct {
	Router     *chi.Mux
	CommonVars *variable.CommonVars
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//var err error

	if h.CommonVars.Name == "" {
		http.Redirect(w, r, "/config", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/show", http.StatusSeeOther)
	}
}
