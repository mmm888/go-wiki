package domain

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mmm888/go-wiki/middleware/markdown"
	"github.com/mmm888/go-wiki/middleware/variable"
)

const (
	defaultFileName = "README.md"
)

type ShowUseCase struct {
}

type ShowInput struct {
	Path       string
	CommonVars *variable.CommonVars
	Markdown   markdown.Markdown
}

type ShowOutput struct {
	Path     string
	Query    string
	Tree     string
	Contents string
}

// ファイルを読み込み、Markdwon to HTML の結果を返す
func (s *ShowUseCase) Get(in *ShowInput) (*ShowOutput, error) {
	root := in.CommonVars.Name
	fpath := filepath.Join(root, in.Path)
	md := in.Markdown

	path, err := checkDirTrav(root, fpath)
	if err != nil {
		return nil, err
	}

	tree, err := dirTree(root, root)
	if err != nil {
		return nil, err
	}

	// root ディレクトリ以外の場所 or 存在しない path のチェック
	if path == "" {
		return &ShowOutput{Path: in.Path, Tree: tree}, nil
	}

	// ディレクトリの場合は defaultFile にアクセス
	if fi, _ := os.Lstat(path); fi.IsDir() {
		path = filepath.Join(path, defaultFileName)

		if _, err := os.Stat(path); err != nil {
			return &ShowOutput{Path: in.Path, Tree: tree}, err
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	out, _ := md.HTMLify(data)

	return &ShowOutput{Path: in.Path, Tree: tree, Contents: out}, nil
}
