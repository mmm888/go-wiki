package bootstrap

import (
	"github.com/mmm888/go-wiki/middleware"
	"github.com/mmm888/go-wiki/middleware/git"
)

func gitSetting(m *middleware.M) {

	root := m.CommonVars.Name
	url := m.CommonVars.Repo

	m.Git = git.NewGit(root, url)
}
