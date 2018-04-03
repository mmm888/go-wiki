package main

import (
	"github.com/gobuffalo/packr"
	"github.com/mmm888/go-wiki/bootstrap"
	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/middleware/assets"
	"github.com/mmm888/go-wiki/middleware/chi"
)

func main() {

	rt := chi.NewRouter()

	// Get bindata
	// TODO: packr を直接使わないと *-packr.go が自動生成されない
	assetses := make(map[string]assets.Assets)
	assetses["public"] = packr.NewBox("./assets/public")
	assetses["templates"] = packr.NewBox("./assets/templates")

	m := &middleware.M{
		Router:   rt,
		Assetses: assetses,
	}

	bootstrap.Start(m)
}
