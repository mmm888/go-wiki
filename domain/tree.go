package domain

import (
	"github.com/mmm888/go-wiki/middleware/variable"
)

type TreeUseCase struct {
}

type TreeInput struct {
	Path       string
	CommonVars *variable.CommonVars
}

type TreeOutput struct {
	Contents string
}

func (s *TreeUseCase) Get(in *TreeInput) (*TreeOutput, error) {
	root := in.CommonVars.Name

	tree, err := dirTree(root, root)
	if err != nil {
		return nil, err
	}

	return &TreeOutput{Contents: tree}, nil
}
