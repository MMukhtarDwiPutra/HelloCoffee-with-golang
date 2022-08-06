package toko

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/go-playground/validator/v10"
)

type Service interface{
	FindAllToko() []model.Toko
	FindOneToko(id int) model.Toko
	FindAllMenuFromToko(id_toko int) []model.Menu
	CreateToko(toko model.Toko) model.Toko
	UpdateToko(t model.Toko, id_toko int) model.Toko
	DeleteToko(id_toko int) model.Toko
}

type service struct{
	repository Repository
	menuRepository menu.Repository
	validate *validator.Validate
}

func NewService(repository Repository, menuRepository menu.Repository) *service{
	validate := validator.New()
	return &service{repository, menuRepository, validate}
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

func (s *service) CreateToko(toko model.Toko) model.Toko{
	err := s.validate.Struct(toko)
	helper.PanicIfError(err)
	s.repository.CreateToko(toko)
	t := s.repository.GetLastToko()
	return t
}

func (s *service) UpdateToko(t model.Toko, id_toko int) model.Toko{
	err := s.validate.Struct(t)
	helper.PanicIfError(err)

	s.repository.UpdateToko(t, id_toko)
	return s.repository.FindOneToko(id_toko)
}

func (s *service) DeleteToko(id_toko int) model.Toko{
	toko := s.repository.FindOneToko(id_toko)
	s.repository.DeleteToko(id_toko)
	return toko
}