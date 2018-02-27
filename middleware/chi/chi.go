package chi

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	return chi.NewRouter()
}

func ListenAndServe(addr string, h http.Handler) error {
	return http.ListenAndServe(addr, h)
}
