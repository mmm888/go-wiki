package middleware

import (
	"github.com/go-chi/chi"
	"github.com/mmm888/go-wiki/middleware/assets"
	"github.com/mmm888/go-wiki/middleware/cron"
	"github.com/mmm888/go-wiki/middleware/git"
	"github.com/mmm888/go-wiki/middleware/markdown"
	"github.com/mmm888/go-wiki/middleware/templates"
	"github.com/mmm888/go-wiki/middleware/variable"
	"github.com/mmm888/go-wiki/middleware/worker"
)

type M struct {
	Router     *chi.Mux
	Assetses   map[string]assets.Assets
	Templates  *templates.Templates
	CommonVars *variable.CommonVars
	Markdown   markdown.Markdown
	JobQueue   *worker.JobQueue
	Git        *git.Git
	Cron       *cron.Cron
}
