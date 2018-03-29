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
	count = 5
	sep   = ":"
	//format = "%h"
	format = "%h" + sep + "%s"
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
func (g *Git) Diff(fpath string, commitHash string) ([]CommitInfo, string, error) {

	path := strings.TrimPrefix(fpath, g.Root+"/")
	if path == g.Root {
		path = "."
	}

	if commitHash == "" {
		// diffのlog一覧を返す
		result, err := g.commitLogList(path)
		if err != nil {
			return nil, "", err
		}

		return result, "", err

	} else {
		// commitHashのgit diff結果を返す
		cmdStr := fmt.Sprintf("-C %s diff %s -- %s", g.Root, commitHash, path)
		cmd := strings.Split(cmdStr, " ")

		result, err := execGit(cmd)
		if err != nil {
			return nil, "", err
		}

		return nil, result, nil
	}
}

type CommitInfo struct {
	Hash    string
	Message string
}

func (g *Git) commitLogList(fpath string) ([]CommitInfo, error) {
	path := strings.TrimPrefix(fpath, g.Root+"/")
	if path == g.Root {
		path = "."
	}

	cmdStr := fmt.Sprintf("-C %s log -%d --pretty=format:\"%s\" -- %s", g.Root, count, format, path)
	cmd := strings.Split(cmdStr, " ")

	// return: <Commit Hash><sep><Commit Message>
	out, err := execGit(cmd)
	if err != nil {
		return nil, err
	}

	list := strings.Split(out, "\n")

	c := count
	if c > len(list) {
		c = len(list)
	}

	result := make([]CommitInfo, c)

	for i := range list {
		line := strings.Trim(list[i], "\"")
		s := strings.Split(line, sep)

		result[i].Hash = s[0]
		result[i].Message = strings.Join(s[1:], "")
	}

	return result, nil
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
