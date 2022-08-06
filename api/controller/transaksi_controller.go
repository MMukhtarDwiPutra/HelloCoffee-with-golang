package controller

import (
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/transaksi"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/exception"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type TransaksiController interface{
	FindAllTransaksi(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	PanicDataNotFound(id_transaksi int)
}

type transaksiController struct{
	transaksiService transaksi.Service
}

func NewTransaksiController(service transaksi.Service) *transaksiController{
	return &transaksiController{service}
}

func (c *transaksiController) FindAllTransaksi(w http.ResponseWriter, r *http.Request){
	transaksiResponses := c.transaksiService.GetAllTransaksiAPI()
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : transaksiResponses,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (c *transaksiController) FindById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_transaksi, _ := strconv.Atoi(params["id"])

	c.PanicDataNotFound(id_transaksi)
	transaksiResponse := c.transaksiService.GetTransaksiAPI(id_transaksi)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : transaksiResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (c *transaksiController) PanicDataNotFound(id_transaksi int){
	transaksiResponse := c.transaksiService.GetTransaksiAPI(id_transaksi)
	if (transaksiResponse == web.TransaksiResponse{}) {
		panic(exception.NewNotFoundError("DATA NOT FOUND"))
	}
}