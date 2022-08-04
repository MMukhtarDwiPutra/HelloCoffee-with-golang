package controller

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type MenuContoller interface{
	FindMenu(w http.ResponseWriter, r *http.Request)
	FindAllMenu(w http.ResponseWriter, r *http.Request)
	CreateMenu(w http.ResponseWriter, r *http.Request)
}

type MenuController struct{
	menuService menu.Service
}

func NewMenuController(service menu.Service) *MenuController{
	return &MenuController{service}
}

func (c *MenuController) FindAllMenu(w http.ResponseWriter, r *http.Request) {
	menuResponses := c.menuService.FindAllMenu()
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *MenuController) FindMenu(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_menu, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)
	menuResponse := c.menuService.FindOneMenu(id_menu)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : menuResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *MenuController) CreateMenu(w http.ResponseWriter, r *http.Request){
	menu := model.Menu{}
	helper.ReadFromRequestBody(r, &menu)

	menuResponse := c.menuService.CreateMenu(menu.Nama_menu, menu.Harga, menu.Deskripsi, menu.Jenis, menu.Id_toko, menu.Foto)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *MenuController) UpdateMenu(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_menu, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)

	menu := model.Menu{}
	helper.ReadFromRequestBody(r, &menu)

	menuResponse := c.menuService.UpdateMenu(menu.Nama_menu, menu.Harga, menu.Deskripsi, menu.Jenis , id_menu)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *MenuController) DeleteMenu(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_menu, err := strconv.Atoi(params["id"])
	helper.PanicIfError(err)

	menuResponse := c.menuService.DeleteMenu(id_menu)
	webResponse := web.WebResponse{
			Code:   200,
		Status: "OK",
		Data:   menuResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}