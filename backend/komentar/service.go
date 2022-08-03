package komentar

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Service interface{
	FindAllKomentar(id_menu int, id_user int) []model.Komentar
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) FindAllKomentar(id_menu int, id_user int) []model.Komentar{
	return s.repository.FindAllKomentar(id_menu, id_user)
}
