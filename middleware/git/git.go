package git

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	count  = 5
	format = "%h"
)

type Git struct {
	Root string
	URL  string
}

func NewGit(root, url string) *Git {
	return &Git{Root: root, URL: url}
}

func (g *Git) Diff(fpath string) (string, error) {
	path := strings.TrimPrefix(fpath, g.Root+"/")
	if path == g.Root {
		path = "."
	}

	out, err := g.commitLogList(path)
	if err != nil {
		return "", err
	}

	cmdStr := fmt.Sprintf("-C %s diff %s -- %s", g.Root, out[len(out)-1], path)
	cmd := strings.Split(cmdStr, " ")

	result, err := execGit(cmd)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (g *Git) commitLogList(fpath string) ([]string, error) {
	path := strings.TrimPrefix(fpath, g.Root+"/")
	if path == g.Root {
		path = "."
	}

	cmdStr := fmt.Sprintf("-C %s log -%d --pretty=format:%s -- %s", g.Root, count, format, path)
	cmd := strings.Split(cmdStr, " ")

	out, err := execGit(cmd)
	if err != nil {
		return nil, err
	}

	return strings.Split(out, "\n"), nil
}

func (g *Git) Init() error {
	// git init
	return nil
}

func (g *Git) Commit() error {
	// git add + git commit
	return nil
}

func (g *Git) Push() error {
	// git push
	return nil
}

func (g *Git) Reset() error {
	// git reset --hard
	return nil
}

func execGit(command []string) (string, error) {
	out, err := exec.Command("git", command...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
