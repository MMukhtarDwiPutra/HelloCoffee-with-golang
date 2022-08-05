package menu

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/go-playground/validator/v10"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
)

type Service interface{
	FindAllMenu() []model.Menu
	FindOneMenu(id_menu int) model.Menu
	CreateMenu(m model.Menu) model.Menu
	DeleteMenu(id_menu int) model.Menu
	UpdateMenu(m model.Menu) model.Menu
}

type service struct{
	repository Repository
	validate *validator.Validate
}

func NewService(repository Repository) *service{
	validate := validator.New()
	return &service{repository, validate}
}

func (s *service) FindAllMenu() []model.Menu{
	return s.repository.FindAllMenu()
}

func (s *service) FindOneMenu(id_menu int) model.Menu{
	return s.repository.FindOneMenu(id_menu)
}

func (s *service) CreateMenu(m model.Menu) model.Menu{
	err := s.validate.Struct(m)
	helper.PanicIfError(err)
	s.repository.CreateMenu(m.Nama_menu, m.Harga, m.Deskripsi, m.Jenis, m.Id_toko, m.Foto)
	
	return s.repository.GetLastMenu()
}

func (s *service) DeleteMenu(id_menu int) model.Menu{
	menu := s.repository.FindOneMenu(id_menu)
	s.repository.DeleteMenu(id_menu)

	return menu
}

func (s *service) UpdateMenu(m model.Menu) model.Menu{
	err := s.validate.Struct(m)
	defer helper.PanicIfError(err)

	s.repository.UpdateMenu(m.Nama_menu, m.Harga, m.Deskripsi, m.Jenis, m.Id_menu)
	
	return s.repository.FindOneMenu(m.Id_menu)
}