package toko

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Service interface{
	FindAllToko() []model.Toko
	FindOneToko(id int) model.Toko
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) FindAllToko() []model.Toko{
	return s.repository.FindAllToko()
}

func (s *service) FindOneToko(id int) model.Toko{
	return s.repository.FindOneToko(id)
}