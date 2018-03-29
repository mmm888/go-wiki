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
	Path         string
	CommitHash   string
	IsCommitHash bool
	CommonVars   *variable.CommonVars
	Git          *git.Git
}

type DiffOutput struct {
	Path         string
	Query        string
	Contents     []string
	IsCommitHash bool
}

func (s *DiffUseCase) Get(in *DiffInput) (*DiffOutput, error) {
	root := in.CommonVars.Name
	fpath := filepath.Join(root, in.Path)
	commitHash := in.CommitHash

	path, err := checkDirTrav(root, fpath)
	if err != nil {
		return nil, err
	}

	out, err := in.Git.Diff(path, commitHash)
	if err != nil {
		return &DiffOutput{Path: in.Path, Contents: []string{}, IsCommitHash: in.IsCommitHash}, err
	}

	// 改行を置換
	if in.IsCommitHash {
		out[0] = strings.Replace(out[0], "\n", "<br>\n", -1)
	}

	return &DiffOutput{Path: in.Path, Contents: out, IsCommitHash: in.IsCommitHash}, nil
}
