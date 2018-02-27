package main

import (
	"flag"
	"log"

	"github.com/mmm888/go-wiki/bootstrap"
	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/middleware/chi"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "address to bind")
	flag.Parse()
}

func main() {
	var err error

	rt := chi.NewRouter()

	m := &middleware.M{
		Router: rt,
	}

	bootstrap.Start(m)

	log.Printf("Start HTTP Server %v", addr)
	err = chi.ListenAndServe(addr, rt)
	if err != nil {
		panic(err)
	}
}
