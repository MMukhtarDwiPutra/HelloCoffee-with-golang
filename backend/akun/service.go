package akun

type Service interface{
	FindOneUser(email string, password string) (int, int)
	CreateAccount(username string, password string, email string, gender string)
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