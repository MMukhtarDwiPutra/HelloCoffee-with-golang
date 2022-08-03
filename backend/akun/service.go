package akun

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
)

type Service interface{
	FindOneUser(email string, password string) (int, int)
	CreateAccount(username string, password string, email string, gender string)
	GetOneAkun(id_user int) model.User
	DeleteAkun(id_user int)
	UpdateAkun(nama string, email string, gender string, id_user int)
	UpdatePassword(pwLama string, pwBaru string, id_user int)
}

type service struct{
	repository Repository
}

func NewService (repository Repository) *service{
	return &service{repository}
}

func (s *service) FindOneUser(email string, password string) (int, int){
	return s.repository.FindOneUser(email, password)
}

func (s *service) CreateAccount(username string, password string, email string, gender string){
	s.repository.CreateAccount(username, password, email, gender)
}

func (s *service) GetOneAkun(id_user int) model.User{
	return s.repository.GetOneAkun(id_user)
}

func (s *service) DeleteAkun(id_user int){
	s.repository.DeleteAkun(id_user)
}

func (s *service) UpdateAkun(nama string, email string, gender string, id_user int){
	s.repository.UpdateAkun(nama, email, gender, id_user)
}

func (s *service) UpdatePassword(pwLama string, pwBaru string, id_user int){
	s.repository.UpdatePassword(pwLama, pwBaru, id_user)
}