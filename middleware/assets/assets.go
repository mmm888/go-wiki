package assets

import "net/http"

type Assets interface {
	Open(string) (http.File, error)
	MustString(string) (string, error)
}
