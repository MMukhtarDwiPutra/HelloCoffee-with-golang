package web

type UserUpdateRequest struct{
	Username string `validate:"required" json:"username"`
	Nama string `validate:"required" json"nama"`
	Password string `validate:"required" json:"password"`
	Email string `validate:"required" json:"email"`
	Gender string `validate:"required" json:"gender"`
	Foto string `validate:"required" json:"foto"`
}