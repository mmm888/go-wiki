package main

import (
	"github.com/mmm888/go-wiki/bootstrap"
	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/middleware/chi"
)

func main() {

	rt := chi.NewRouter()

	m := &middleware.M{
		Router: rt,
	}

	bootstrap.Start(m)
}
