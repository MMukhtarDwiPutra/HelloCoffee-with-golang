package keranjang

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/transaksi"
	"github.com/go-playground/validator/v10"
)

type Service interface{
	GetAllKeranjang(id_user int) ([]model.Keranjang, int)
	AddKeranjang(id_menu int, qty int, id_user int)
	AddKeranjangAPI(k web.CreateKeranjangRequest)
	DeleteKeranjang(id_user int)
	CheckoutProcess(id_user int, fullName string, email string, address string, city string, zip string, state string)
	GetLastKeranjang() web.KeranjangResponse
	GetKeranjangAPI(id_keranjang int) web.KeranjangResponse
	DeleteKeranjangAPI(id_keranjang int)
	FindAllKeranjangAPI() []web.KeranjangResponse
	UpdateKeranjangAPI(k web.CreateKeranjangRequest, id_keranjang int)
}

type service struct{
	repository Repository
	transaksiRepository transaksi.Repository
	validate *validator.Validate
}

func NewService(repository Repository, keranjangRepository transaksi.Repository) *service{
	validate := validator.New()
	return &service{repository, keranjangRepository, validate}
}

func (s *service) GetAllKeranjang(id_user int) ([]model.Keranjang, int){
	return s.repository.GetAllKeranjang(id_user)
}

func (s *service) AddKeranjang(id_menu int, qty int, id_user int){

	s.repository.AddKeranjang(id_menu, qty, id_user)
}

func (s *service) AddKeranjangAPI(k web.CreateKeranjangRequest){
	err := s.validate.Struct(k)
	helper.PanicIfError(err)

	s.repository.AddKeranjang(k.IdMenu, k.Qty, k.IdUser)
}

func (s *service) DeleteKeranjang(id_user int){
	s.repository.DeleteAllKeranjang(id_user)
}

func (s *service) CheckoutProcess(id_user int, fullName string, email string, address string, city string, zip string, state string){
	s.transaksiRepository.AddTransaksi(fullName, email, address, city, zip, state)
	s.repository.AddTransaksi(id_user)
}

func (s *service) GetLastKeranjang() web.KeranjangResponse{
	return s.repository.GetLastKeranjang()
}

func (s *service) GetKeranjangAPI(id_keranjang int) web.KeranjangResponse{
	return s.repository.GetKeranjangAPI(id_keranjang)
}

func (s *service) UpdateKeranjangAPI(k web.CreateKeranjangRequest, id_keranjang int){
	err := s.validate.Struct(k)
	helper.PanicIfError(err)

	s.repository.UpdateKeranjangAPI(k, id_keranjang)
}

func (s *service) DeleteKeranjangAPI(id_keranjang int){
	s.repository.DeleteKeranjangAPI(id_keranjang)
}

func (s *service) FindAllKeranjangAPI() []web.KeranjangResponse{
	return s.repository.FindAllKeranjangAPI()
}