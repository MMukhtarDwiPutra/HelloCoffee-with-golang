package controller

import(
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/keranjang"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model/web"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/exception"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type KeranjangController interface{
	CreateKeranjang(w http.ResponseWriter, r *http.Request)
	PanicDataNotFound(id_keranjang int)
	FindById(w http.ResponseWriter, r *http.Request)
	UpdateKeranjang(w http.ResponseWriter, r *http.Request)
}

type keranjangController struct{
	keranjangService keranjang.Service
}

func NewKeranjangController(service keranjang.Service) *keranjangController{
	return &keranjangController{service}
}

func (c *keranjangController) CreateKeranjang(w http.ResponseWriter, r *http.Request){
	keranjang := web.CreateKeranjangRequest{}
	helper.ReadFromRequestBody(r, &keranjang)

	c.keranjangService.AddKeranjangAPI(keranjang)
	keranjangResponse := c.keranjangService.GetLastKeranjang()
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : keranjangResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *keranjangController) FindById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_keranjang, _ := strconv.Atoi(params["id"])

	c.PanicDataNotFound(id_keranjang)
	keranjangResponse := c.keranjangService.GetKeranjangAPI(id_keranjang)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : keranjangResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *keranjangController) UpdateKeranjang(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_keranjang, _ := strconv.Atoi(params["id"])

	keranjang := web.CreateKeranjangRequest{}
	helper.ReadFromRequestBody(r, &keranjang)

	c.PanicDataNotFound(id_keranjang)
	c.keranjangService.UpdateKeranjangAPI(keranjang, id_keranjang)
	keranjangResponse := c.keranjangService.GetKeranjangAPI(id_keranjang)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : keranjangResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *keranjangController) DeleteKeranjang(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id_keranjang, _ := strconv.Atoi(params["id"])

	c.PanicDataNotFound(id_keranjang)
	keranjangResponse := c.keranjangService.GetKeranjangAPI(id_keranjang)

	c.keranjangService.DeleteKeranjangAPI(id_keranjang)
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : keranjangResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *keranjangController) FindAllKeranjang(w http.ResponseWriter, r *http.Request){
	keranjangResponses := c.keranjangService.FindAllKeranjangAPI()
	webResponse := web.WebResponse{
		Code : 200,
		Status : "OK",
		Data : keranjangResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (c *keranjangController) PanicDataNotFound(id_keranjang int){
	keranjangResponse := c.keranjangService.GetKeranjangAPI(id_keranjang)
	if (keranjangResponse == web.KeranjangResponse{}) {
		panic(exception.NewNotFoundError("DATA NOT FOUND"))
	}
}