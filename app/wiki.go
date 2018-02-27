package app

import wiki "github.com/mmm888/go-wiki/domain"

type ShowService struct {
	Input *wiki.ShowInput
	Info  *wiki.ShowUseCase
}

type EditService struct {
	Info *wiki.EditUseCase
}

type ConfigService struct {
	Input *wiki.ConfigInput
	Info  *wiki.ConfigUseCase
}

type DiffService struct {
	Input *wiki.DiffInput
	Info  *wiki.DiffUseCase
}
