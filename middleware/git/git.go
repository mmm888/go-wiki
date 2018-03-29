package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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

// commitHashに値が入っていない場合は一覧を返し、
// 入っている場合はdiffの内容をスライスの先頭に入れて返す
func (g *Git) Diff(fpath string, commitHash string) ([]string, error) {

	var err error

	path := strings.TrimPrefix(fpath, g.Root+"/")
	if path == g.Root {
		path = "."
	}

	var result []string

	if commitHash == "" {
		// diffのlog一覧を返す
		result, err = g.commitLogList(path)
		if err != nil {
			return nil, err
		}

	} else {
		// commitHashのgit diff結果を返す
		cmdStr := fmt.Sprintf("-C %s diff %s -- %s", g.Root, commitHash, path)
		cmd := strings.Split(cmdStr, " ")

		out, err := execGit(cmd)
		if err != nil {
			return nil, err
		}

		result = []string{out}
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

func (g *Git) Init() (string, error) {

	// check g.Root/.git
	dotGitPath := filepath.Join(g.Root, ".git")
	if _, err := os.Stat(dotGitPath); err == nil {
		return "Nothing", nil
	}

	// mkdir g.Root
	if _, err := os.Stat(g.Root); err != nil {
		if err := os.MkdirAll(g.Root, 0755); err != nil {
			return "", err
		}
	}

	cmdStr := fmt.Sprintf("-C %s init", g.Root)
	cmd := strings.Split(cmdStr, " ")

	result, err := execGit(cmd)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (g *Git) Commit() (string, error) {
	// git add
	cmdStr := fmt.Sprintf("-C %s add -A", g.Root)
	cmd := strings.Split(cmdStr, " ")

	result, err := execGit(cmd)
	if err != nil {
		return "", err
	}

	// 現在時間取得
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst)
	nowStr := now.Format("2006-01-02")

	// git commit
	cmdStr = fmt.Sprintf("-C %s commit -m %s", g.Root, nowStr)
	cmd = strings.Split(cmdStr, " ")

	out, err := execGit(cmd)
	if err != nil {
		return "", err
	}
	result += out

	return result, nil
}

func (g *Git) Push() (string, error) {
	// git push
	return "", nil
}

func (g *Git) Reset() (string, error) {
	// git reset --hard
	return "", nil
}

func execGit(command []string) (string, error) {
	out, err := exec.Command("git", command...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
