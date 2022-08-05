package controller

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/exception"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type UserController interface{
	CreateUser(w http.ResponseWriter, r *http.Request)
	FindAllUser(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	PanicDataNotFound(id_user int) model.Menu
}

type userController struct{
	userService akun.Service
}

func NewUserController(service akun.Service) *userController{
	return &userController{service}
}

func (c *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	helper.ReadFromRequestBody(r, &user)

	userResponse := c.userService.CreateAccount(user)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *userController) FindAllUser(w http.ResponseWriter, r *http.Request){
	userResponses := c.userService.FindAllUser()

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : userResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *userController) FindById(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
 	id_user, _ := strconv.Atoi(id)
	userResponse :=  c.PanicDataNotFound(id_user)

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *userController) UpdateUser(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	id_user, _ := strconv.Atoi(id)
	c.PanicDataNotFound(id_user)

	userRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(r, &userRequest)
	userResponse := c.userService.UpdateAkunAPI(userRequest, id_user)

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *userController) DeleteUser(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	id_user, _ := strconv.Atoi(id)

	c.PanicDataNotFound(id_user)
	userResponse := c.userService.DeleteAkun(id_user)

	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *userController) PanicDataNotFound(id_user int) model.User{
	userResponse := c.userService.GetOneAkun(id_user)
	if (userResponse == model.User{}) {
		panic(exception.NewNotFoundError("DATA NOT FOUND"))
	}

	return userResponse
}