package middleware

import (
	"html/template"

	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type M struct {
	Router     *chi.Mux
	Template   *template.Template
	CommonVars *variable.CommonVars
}
