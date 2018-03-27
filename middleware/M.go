package middleware

import (
	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/middleware/markdown"
	"github.com/mmm888/go-wiki/middleware/templates"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type M struct {
	Router     *chi.Mux
	Templates  *templates.Templates
	CommonVars *variable.CommonVars
	Markdown   markdown.Markdown
}
