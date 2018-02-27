package domain

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mmm888/go-wiki/middleware/variable"
)

type ConfigUseCase struct {
}

type ConfigInput struct {
	Name       string
	Repo       string
	Path       string
	CommonVars *variable.CommonVars
}

type ConfigOutput struct {
	Path  string
	Query string
	Name  string
	Repo  string
}

func (u *ConfigUseCase) Get(in *ConfigInput) (*ConfigOutput, error) {
	return &ConfigOutput{
		Path: in.Path,
		Name: in.CommonVars.Name,
		Repo: in.CommonVars.Repo,
	}, nil
}

func (u *ConfigUseCase) Post(in *ConfigInput) error {
	v := in.CommonVars

	v.Name = in.Name
	v.Repo = in.Repo

	data, err := json.Marshal(in.CommonVars)
	if err != nil {
		return err
	}

	f, err := os.Create(v.ConfigPath)
	if err != nil {
		return err
	}
	fmt.Fprint(f, string(data))

	return nil
}
