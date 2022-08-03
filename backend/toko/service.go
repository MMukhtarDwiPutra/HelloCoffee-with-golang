package toko

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
)

type Service interface{
	FindAllToko() []model.Toko
	FindOneToko(id int) model.Toko
	FindAllMenuFromToko(id_toko int) []model.Menu
}

type service struct{
	repository Repository
	menuRepository menu.Repository
}

func NewService(repository Repository, menuRepository menu.Repository) *service{
	return &service{repository, menuRepository}
}

func (s *service) FindAllToko() []model.Toko{
	return s.repository.FindAllToko()
}

func (s *service) FindOneToko(id int) model.Toko{
	return s.repository.FindOneToko(id)
}

func (s *service) FindAllMenuFromToko(id_toko int) []model.Menu{
	return s.menuRepository.FindAllMenuFromToko(id_toko)
}