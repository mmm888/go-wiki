package bootstrap

import (
	"net/http"
	"path"

	"github.com/go-chi/chi"
	chiMiddle "github.com/go-chi/chi/middleware"
	"github.com/mmm888/go-wiki/app"
	wiki "github.com/mmm888/go-wiki/domain"
	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/web"
)

func registerRoute(m *middleware.M) {
	r := m.Router

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
		fileServer(r, "/css", "./public/css")
		fileServer(r, "/js", "./public/js")

		// 最後に削除
		fileServer(r, "/html", "./public/html")
	}
}

// TODO: https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
func fileServer(r *chi.Mux, pattern, filepath string) {
	fs := http.StripPrefix(pattern, http.FileServer(http.Dir(filepath)))
	pattern = path.Join(pattern, "*")

	r.Get(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
