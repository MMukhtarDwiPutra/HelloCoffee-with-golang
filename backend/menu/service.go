package menu

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Service interface{
	FindAllMenu() []model.Menu
	FindOneMenu(id_menu int) model.Menu
	CreateMenu(nama string, harga int, deskripsi string, jenis string, id_toko int, foto_kopi string) model.Menu
	DeleteMenu(id_menu int) model.Menu
	UpdateMenu(nama string, harga int, deskripsi string, jenis string, id_menu int) model.Menu
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) FindAllMenu() []model.Menu{
	return s.repository.FindAllMenu()
}

func (s *service) FindOneMenu(id_menu int) model.Menu{
	return s.repository.FindOneMenu(id_menu)
}

func (s *service) CreateMenu(nama string, harga int, deskripsi string, jenis string, id_toko int, foto_kopi string) model.Menu{
	s.repository.CreateMenu(nama, harga, deskripsi, jenis, id_toko, foto_kopi)
	
	return s.repository.GetLastMenu()
}

func (s *service) DeleteMenu(id_menu int) model.Menu{
	menu := s.repository.FindOneMenu(id_menu)
	s.repository.DeleteMenu(id_menu)

	return menu
}

func (s *service) UpdateMenu(nama string, harga int, deskripsi string, jenis string, id_menu int) model.Menu{
	s.repository.UpdateMenu(nama, harga, deskripsi, jenis, id_menu)
	
	return s.repository.FindOneMenu(id_menu)
}