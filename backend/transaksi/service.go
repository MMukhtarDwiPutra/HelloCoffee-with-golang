package transaksi

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
)

type Service interface{
	GetAllTransaksi(id_user int) []model.Transaksi
	GetTransaksiToko(id_toko int) []model.Transaksi
	UpdateStatusTransaksi(status string, id_transaksi int)
	GetTransaksiAPI(id_transaksi int) web.TransaksiResponse
	GetAllTransaksiAPI() []web.TransaksiResponse
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) GetAllTransaksi(id_user int) []model.Transaksi{
	return s.repository.GetAllTransaksi(id_user)
}

func (s *service) GetTransaksiToko(id_toko int) []model.Transaksi{
	return s.repository.GetTransaksiToko(id_toko)
}

func (s *service) UpdateStatusTransaksi(status string, id_transaksi int){
	s.repository.UpdateStatusTransaksi(status, id_transaksi)
}

func (s *service) GetTransaksiAPI(id_transaksi int) web.TransaksiResponse{
	return s.repository.GetTransaksiAPI(id_transaksi)
}

func (s *service) GetAllTransaksiAPI() []web.TransaksiResponse{
	return s.repository.GetAllTransaksiAPI()
}