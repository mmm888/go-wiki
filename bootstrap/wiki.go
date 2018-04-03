package bootstrap

import (
	"net/http"

	chiMiddle "github.com/go-chi/chi/middleware"
	"github.com/mmm888/go-wiki/app"
	wiki "github.com/mmm888/go-wiki/domain"
	"github.com/mmm888/go-wiki/job"
	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/web"
)

func registerRoute(m *middleware.M) {
	r := m.Router
	jq := m.JobQueue

	// logger
	r.Use(chiMiddle.Logger)

	r.Method("GET", "/", &web.RootHandler{
		Router:     m.Router,
		CommonVars: m.CommonVars,
	})

	{
		r.Method("GET", "/config", &web.ConfigGetHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Templates:  m.Templates,
			Service:    &app.ConfigService{Info: &wiki.ConfigUseCase{}},
		})
		r.Method("POST", "/config", &web.ConfigPostHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Service:    &app.ConfigService{Info: &wiki.ConfigUseCase{}},
		})
	}

	{
		r.Method("GET", "/show", &web.ShowHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Templates:  m.Templates,
			Markdown:   m.Markdown,
			Service:    &app.ShowService{Info: &wiki.ShowUseCase{}},
		})
	}

	{
		r.Method("GET", "/tree", &web.TreeHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Templates:  m.Templates,
			Service:    &app.TreeService{Info: &wiki.TreeUseCase{}},
		})
	}

	{
		r.Method("GET", "/edit", &web.EditGetHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Templates:  m.Templates,
			Service:    &app.EditService{Info: &wiki.EditUseCase{}},
		})
		r.Method("POST", "/edit", &web.EditPostHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Service:    &app.EditService{Info: &wiki.EditUseCase{}},
		})
	}

	{
		r.Method("GET", "/diff", &web.DiffHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Templates:  m.Templates,
			Service:    &app.DiffService{Info: &wiki.DiffUseCase{}},
		})

		r.Method("GET", "/diff/{hash}", &web.DiffHandler{
			Router:     m.Router,
			CommonVars: m.CommonVars,
			Templates:  m.Templates,
			Service:    &app.DiffService{Info: &wiki.DiffUseCase{}},
		})
	}

	// static file
	{
		publicAssets := m.Assetses["public"]
		r.Method("GET", "/css/*", http.FileServer(publicAssets))
		r.Method("GET", "/js/*", http.FileServer(publicAssets))

		// TODO: テスト用, 最後に削除
		r.Method("GET", "/html/*", http.FileServer(publicAssets))
	}

	// worker
	{
		j := job.GitCommitJob{
			Git: m.Git,
		}
		jq.Route("git/commit", j)
	}
}
