package domain

import (
	"path/filepath"
	"strings"

	"github.com/mmm888/go-wiki/middleware/git"
	"github.com/mmm888/go-wiki/middleware/variable"
)

type DiffUseCase struct {
}

type DiffInput struct {
	Path       string
	CommonVars *variable.CommonVars
	Git        *git.Git
}

type DiffOutput struct {
	Path     string
	Query    string
	Contents string
}

func (s *DiffUseCase) Get(in *DiffInput) (*DiffOutput, error) {
	root := in.CommonVars.Name
	fpath := filepath.Join(root, in.Path)

	path, err := checkDirTrav(root, fpath)
	if err != nil {
		return nil, err
	}

	out, err := in.Git.Diff(path)
	if err != nil {
		return nil, err
	}

	// 改行を置換
	out = strings.Replace(out, "\n", "<br>\n", -1)

	return &DiffOutput{Path: in.Path, Contents: out}, nil
}
