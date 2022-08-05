package akun

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
	"github.com/go-playground/validator/v10"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
)

type Service interface{
	FindOneUser(email string, password string) (int, int)
	CreateAccount(user model.User) model.User
	GetOneAkun(id_user int) model.User
	DeleteAkun(id_user int) model.User
	UpdateAkun(user model.User)
	UpdatePassword(pwLama string, pwBaru string, id_user int)
	FindAllUser() []model.User
	UpdateAkunAPI(user web.UserUpdateRequest, id_user int) model.User
}

type service struct{
	repository Repository
	validate *validator.Validate
}

func NewService (repository Repository) *service{
	validate := validator.New()
	return &service{repository, validate}
}

func (s *service) FindAllUser() []model.User{
	return s.repository.FindAllUser()
}

func (s *service) FindOneUser(email string, password string) (int, int){
	return s.repository.FindOneUser(email, password)
}

func (s *service) CreateAccount(user model.User) model.User{
	err := s.validate.Struct(user)
	helper.PanicIfError(err)
	s.repository.CreateAccount(user.Username, user.Password, user.Email, user.Nama, user.Gender, user.Foto, user.Id_toko)

	return s.repository.GetLastUser()
}

func (s *service) GetOneAkun(id_user int) model.User{
	user := s.repository.GetOneAkun(id_user)

	return user
}

func (s *service) DeleteAkun(id_user int) model.User{
	user := s.repository.GetOneAkun(id_user)
	s.repository.DeleteAkun(id_user)

	return user
}

func (s *service) UpdateAkun(user model.User){
	s.repository.UpdateAkun(user.Nama, user.Email, user.Gender, user.Id_user)
}

func (s *service) UpdateAkunAPI(user web.UserUpdateRequest, id_user int) model.User{
	err := s.validate.Struct(user)
	helper.PanicIfError(err)
	s.repository.UpdateAkunAPI(user.Username, user.Password, user.Email, user.Gender, user.Nama, user.Foto, id_user)

	return s.repository.GetOneAkun(id_user)
}

func (s *service) UpdatePassword(pwLama string, pwBaru string, id_user int){
	s.repository.UpdatePassword(pwLama, pwBaru, id_user)
}