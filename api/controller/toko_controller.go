package controller

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/toko"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/exception"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type TokoController interface{
	CreateToko(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAllToko(w http.ResponseWriter, r *http.Request)
	UpdateToko(w http.ResponseWriter, r *http.Request)
	DeleteToko(w http.ResponseWriter, r *http.Request)
	PanicDataNotFound(id_toko int) model.Menu
}

type tokoController struct{
	tokoService toko.Service
}

func NewTokoController(service toko.Service) *tokoController{
	return &tokoController{service}
}

func (c *tokoController) CreateToko(w http.ResponseWriter, r *http.Request){
	toko := model.Toko{}
	helper.ReadFromRequestBody(r, &toko)

	tokoResponse := c.tokoService.CreateToko(toko)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : tokoResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (c *tokoController) FindById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_toko, _ := strconv.Atoi(params["id"])

	c.PanicDataNotFound(id_toko)
	tokoResponse := c.tokoService.FindOneToko(id_toko)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : tokoResponse,
	}
	helper.WriteToResponseBody(w, webResponse) 
}

func (c *tokoController) FindAllToko(w http.ResponseWriter, r *http.Request){
	tokoResponse := c.tokoService.FindAllToko()
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : tokoResponse,
	}
	helper.WriteToResponseBody(w, webResponse) 
}

func (c *tokoController) UpdateToko(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_toko, _ := strconv.Atoi(params["id"])

	c.PanicDataNotFound(id_toko)
	toko := model.Toko{}
	helper.ReadFromRequestBody(r, &toko)

	tokoResponse := c.tokoService.UpdateToko(toko, id_toko)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : tokoResponse,
	}
	helper.WriteToResponseBody(w, webResponse) 
}

func (c *tokoController) DeleteToko(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_toko, _ := strconv.Atoi(params["id"])
	c.PanicDataNotFound(id_toko)

	tokoResponse := c.tokoService.DeleteToko(id_toko)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : tokoResponse,
	}
	helper.WriteToResponseBody(w, webResponse) 
}

func (c *tokoController) PanicDataNotFound(id_toko int) model.Toko{
	tokoResponse := c.tokoService.FindOneToko(id_toko)
	if (tokoResponse == model.Toko{}) {
		panic(exception.NewNotFoundError("DATA NOT FOUND"))
	}
	return tokoResponse
}