package templates

import (
	"errors"
	"io"
	"text/template"

	"github.com/mmm888/go-wiki/middleware/assets"
)

type Templates struct {
	root    string
	tmpls   map[string]*template.Template
	funcMap template.FuncMap
	// TODO: 依存するのが良くないかもしれない
	assets assets.Assets
}

/*
func (t *Templates) Route(r string, fs ...string) error {
	var err error

	f := fs[0]

	var files []string
	for i := range fs {
		path := filepath.Join(t.root, fs[i])
		files = append(files, path)
	}

	t.tmpls[r], err = template.New(f).Funcs(t.funcMap).ParseFiles(files...)
	if err != nil {
		return err
	}

	return nil
}
*/

func (t *Templates) Route(r string, fs ...string) error {
	var err error

	f := fs[0]

	t.tmpls[r] = template.New(f).Funcs(t.funcMap)
	if err != nil {
		return err
	}

	for i := range fs {
		//path := filepath.Join(t.root, fs[i])
		path := fs[i]

		text, err := t.assets.MustString(path)
		if err != nil {
			return err
		}

		t.tmpls[r], err = t.tmpls[r].Parse(text)
		if err != nil {
			return err
		}
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
func NewTemplates(root string, assets assets.Assets) *Templates {
	return &Templates{
		root:    root,
		tmpls:   make(map[string]*template.Template),
		funcMap: template.FuncMap{},
		assets:  assets,
	}
}
