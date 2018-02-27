package domain

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ディレクトリトラバーサルのチェック
func checkDirTrav(root, fpath string) (string, error) {
	dir := strings.TrimPrefix(fpath, root+"/")
	if strings.HasPrefix(dir, ".git") {
		return "", errors.New(".git ディレクトリです")
	}

	var p string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fpath == path {
			p = path
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return p, nil
}

// ディレクトリツリーの HTML を作成
func dirTree(originRoot, root string) (string, error) {
	var tree string
	tree += fmt.Sprintln("<ul>")

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return "", nil
	}

	for _, file := range files {
		name := file.Name()
		if name == ".git" {
			continue
		}

		path := filepath.Join(root, name)
		tree += fmt.Sprintf("<li><a href=\"/show?path=%s\">%s</a></li>\n", strings.TrimPrefix(path, originRoot+"/"), name)
		if file.IsDir() {
			dir, err := dirTree(originRoot, path)
			if err != nil {
				return "", err
			}

			tree += dir
		}
	}

	tree += fmt.Sprintln("</ul>")

	return tree, nil
}
