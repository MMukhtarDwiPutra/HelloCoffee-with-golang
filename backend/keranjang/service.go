package keranjang

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/transaksi"
)

type Service interface{
	GetAllKeranjang(id_user int) ([]model.Keranjang, int)
	AddKeranjang(id_menu int, qty int, id_user int)
	DeleteKeranjang(id_user int)
	CheckoutProcess(id_user int, fullName string, email string, address string, city string, zip string, state string)
}

type service struct{
	repository Repository
	transaksiRepository transaksi.Repository
}

func NewService(repository Repository, keranjangRepository transaksi.Repository) *service{
	return &service{repository, keranjangRepository}
}

func (s *service) GetAllKeranjang(id_user int) ([]model.Keranjang, int){
	return s.repository.GetAllKeranjang(id_user)
}

func (s *service) AddKeranjang(id_menu int, qty int, id_user int){
	s.repository.AddKeranjang(id_menu, qty, id_user)
}

func (s *service) DeleteKeranjang(id_user int){
	s.repository.DeleteAllKeranjang(id_user)
}

func (s *service) CheckoutProcess(id_user int, fullName string, email string, address string, city string, zip string, state string){
	s.transaksiRepository.AddTransaksi(fullName, email, address, city, zip, state)
	s.repository.AddTransaksi(id_user)
}