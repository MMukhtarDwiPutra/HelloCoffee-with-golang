package main

import (
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"

	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/toko"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/komentar"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/handler"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/transaksi"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/keranjang"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/app"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/helper"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/application"
)

func main(){
	db := app.NewDB()

	transaksiRepository := transaksi.NewRepository(db)
	transaksiService := transaksi.NewService(transaksiRepository)

	komentarRepository := komentar.NewRepository(db)
	komentarService := komentar.NewService(komentarRepository)

	akunRepository := akun.NewRepository(db)
	akunService := akun.NewService(akunRepository)
	akunHandler := handler.NewAkunHandler(akunService, transaksiService, komentarService)

	menuRepository := menu.NewRepository(db)
	menuService := menu.NewService(menuRepository)
	menuHandler := handler.NewMenuHandler(menuService, komentarService)
	
	tokoRepository := toko.NewRepository(db)
	tokoService := toko.NewService(tokoRepository, menuRepository)
	tokoHandler := handler.NewTokoHandler(tokoService)

	keranjangRepository := keranjang.NewRepository(db)
	keranjangService := keranjang.NewService(keranjangRepository, transaksiRepository)
	keranjangHandler := handler.NewKeranjangHandler(keranjangService)

	router := application.NewRouter(akunHandler, tokoHandler, menuHandler, keranjangHandler)
	router = app.AddRouterAPI(router, menuService, akunService, tokoService, keranjangService, transaksiService)
    
	log.Println("SERVER is running at port 8080")

	err := http.ListenAndServe(":8080", router)
	helper.PanicIfError(err)
}