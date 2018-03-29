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
	DiffList     []git.CommitInfo
	DiffInfo     string
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

	list, info, err := in.Git.Diff(path, commitHash)
	if err != nil {
		return &DiffOutput{Path: in.Path, IsCommitHash: in.IsCommitHash}, err
	}

	// 改行を置換
	if in.IsCommitHash {
		info = strings.Replace(info, "\n", "<br>\n", -1)
	}

	return &DiffOutput{Path: in.Path, DiffList: list, DiffInfo: info, IsCommitHash: in.IsCommitHash}, nil
}
