package domain

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mmm888/go-wiki/middleware/variable"
)

const (
	typeWriteFile       = ""
	typeCreateFile      = "type_file"
	typeCreateDirectory = "type_dir"

	defaultDirectoryPermission = 0755
	defaultFilePermission      = 0644
)

type EditUseCase struct {
}

type EditInput struct {
	Path       string
	Name       string
	FileType   string
	Contents   string
	CommonVars *variable.CommonVars
}

type EditOutput struct {
	Path     string
	Query    string
	File     string
	Contents string
	IsDir    bool
}

func (u *EditUseCase) Get(in *EditInput) (*EditOutput, error) {
	root := in.CommonVars.Name
	fpath := filepath.Join(root, in.Path)

	path, err := checkDirTrav(root, fpath)
	if err != nil {
		return nil, err
	}

	var isDir bool
	if fi, _ := os.Lstat(path); fi.IsDir() {
		isDir = true

		path = filepath.Join(path, defaultFileName)
		if _, err := os.Stat(path); err != nil {
			return &EditOutput{Path: in.Path, File: defaultFileName, IsDir: isDir}, err
		}
	}

	filename := filepath.Base(path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return &EditOutput{Path: in.Path, File: filename, Contents: string(data), IsDir: isDir}, nil
}

func (u *EditUseCase) Post(in *EditInput) error {
	root := in.CommonVars.Name
	fpath := filepath.Join(root, in.Path)

	switch in.FileType {
	case typeWriteFile:
		writePath, err := checkDirTrav(root, fpath)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(writePath, []byte(in.Contents), defaultFilePermission); err != nil {
			return err
		}

	case typeCreateFile:
		dir, err := checkDirTrav(root, fpath)
		if err != nil {
			return err
		}

		createPath := filepath.Join(dir, in.Name)
		if _, err := os.Stat(createPath); err == nil {
			return errors.New("すでに存在します")
		}

		_, err = os.Create(createPath)
		if err != nil {
			return err
		}

	case typeCreateDirectory:
		dir, err := checkDirTrav(root, fpath)
		if err != nil {
			return err
		}

		createPath := filepath.Join(dir, in.Name)
		if _, err := os.Stat(createPath); err == nil {
			return errors.New("すでに存在します")
		}

		err = os.Mkdir(createPath, defaultDirectoryPermission)
		if err != nil {
			return err
		}
	}

	return nil
}
