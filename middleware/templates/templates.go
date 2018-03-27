package templates

import (
	"errors"
	"io"
	"log"
	"path/filepath"
	"text/template"
)

type Templates struct {
	root    string
	tmpls   map[string]*template.Template
	funcMap template.FuncMap
}

func (t *Templates) Route(r, f string) error {
	var err error

	path := filepath.Join(t.root, f)

	log.Print("template.New(", f, ").Funcs(", t.funcMap, ").ParseFiles(", path, ")")

	t.tmpls[r], err = template.New(f).Funcs(t.funcMap).ParseFiles(path)
	if err != nil {
		return err
	}

	return nil
}

func (t *Templates) Render(r string, wr io.Writer, data interface{}) error {
	var tmpl *template.Template
	var ok bool

	if tmpl, ok = t.tmpls[r]; !ok {
		return errors.New("Not data")
	}

	if err := tmpl.Execute(wr, data); err != nil {
		return err
	}

	return nil
}

// TODO: funcMap{} 初期化 or 削除
func NewTemplates(root string) *Templates {
	return &Templates{
		root:    root,
		tmpls:   make(map[string]*template.Template),
		funcMap: template.FuncMap{},
	}
}
