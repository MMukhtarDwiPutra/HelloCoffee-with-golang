package komentar

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Service interface{
	FindAllKomentar(id_menu int, id_user int) []model.Komentar
	AddKomentar(komentar string, id_user int, id_menu int)
	DeleteKomentar(id_komentar int)
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

func (s *service) AddKomentar(komentar string, id_user int, id_menu int){
	s.repository.AddKomentar(komentar, id_user, id_menu)
}

func (s *service) DeleteKomentar(id_komentar int){
	s.repository.DeleteKomentar(id_komentar)
}