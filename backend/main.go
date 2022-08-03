package main

import (
	"database/sql"
	// "embed"
	// "io/fs"
	"log"
	"net/http"

	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/toko"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/komentar"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/handler"
	// "github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"
)

func main(){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/hellocoffee")
	if err != nil{
		log.Fatal(err)
	}

	akunRepository := akun.NewRepository(db)
	akunService := akun.NewService(akunRepository)
	akunHandler := handler.NewAkunHandler(akunService)

	tokoRepository := toko.NewRepository(db)
	tokoService := toko.NewService(tokoRepository)
	tokoHandler := handler.NewTokoHandler(tokoService)

	komentarRepository := komentar.NewRepository(db)
	komentarService := komentar.NewService(komentarRepository)

	menuRepository := menu.NewRepository(db)
	menuService := menu.NewService(menuRepository)
	menuHandler := handler.NewMenuHandler(menuService, komentarService)

	router := mux.NewRouter()

	var dataAkun akun.DataAkun
	dataAkun = akun.NewDataAkun()

	router.HandleFunc("/akun/login", akunHandler.LoginProcess)
	router.HandleFunc("/akun/register", akunHandler.RegisterHandler)
	router.HandleFunc("/akun/registration", akunHandler.RegisterNewAccount)
	router.HandleFunc("/", akunHandler.LoginHandler)

	router.HandleFunc("/home", tokoHandler.HomeHandler)
	router.HandleFunc("/toko/", tokoHandler.DetailTokoHandler)

	router.HandleFunc("/menu", menuHandler.MenuHandler)
	router.HandleFunc("/menu/detail/", menuHandler.DetailMenu)
	router.HandleFunc("/menu/edit/", menuHandler.EditMenuHandler)
	router.HandleFunc("/menu/tambahMenu", menuHandler.TambahMenuHandler)
	router.HandleFunc("/menu/tambahMenu/process", menuHandler.TambahMenuProcess)
	router.HandleFunc("/menu/hapus/", menuHandler.DeleteMenu)
	router.HandleFunc("/menu/edit/process/", menuHandler.EditMenuProcess)

	router.HandleFunc("/akun/pengaturan/", dataAkun.SettingHandler)
	router.HandleFunc("/logout", dataAkun.LogoutHandler)
	router.HandleFunc("/akun/pengaturan/deleteAkun/", dataAkun.DeleteAkun)
	router.HandleFunc("/akun/pengaturan/edit/", dataAkun.EditAkunHandler)
	router.HandleFunc("/akun/pengaturan/edit/process/",dataAkun.EditAkun)
	router.HandleFunc("/akun/pengaturan/edit/password/", dataAkun.EditPasswordHandler)
	router.HandleFunc("/akun/pengaturan/edit/password/process/",dataAkun.EditPassword)
	router.HandleFunc("/komentar/tambahKomentar/", dataAkun.TambahKomentar)
	router.HandleFunc("/komentar/hapusKomentar/", dataAkun.HapusKomentar)
	router.HandleFunc("/keranjang/", dataAkun.KeranjangHandler)
	router.HandleFunc("/keranjang/tambahKeranjang/", dataAkun.TambahKeranjang)
	router.HandleFunc("/keranjang/hapusSemua/", dataAkun.HapusKeranjang)
	router.HandleFunc("/keranjang/checkout/", dataAkun.CheckoutHandler)
	router.HandleFunc("/checkout/process/", dataAkun.CheckoutProcess)
	router.HandleFunc("/home/toko", dataAkun.HomeTokoHandler)
	router.HandleFunc("/transaksi/", dataAkun.TransaksiHandler)
	router.HandleFunc("/transaksi/process/", dataAkun.ProcessTransaksi)
	router.HandleFunc("/keranjang/checkOutNow/", dataAkun.CheckoutNowHandler)

	path, _ := os.Getwd()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(path+`\backend\assets`))))
	router.PathPrefix("/static/img/").Handler(http.StripPrefix("/static/img/", http.FileServer(http.Dir(path+`\backend\assets\img`))))
    
	log.Println("SERVER is running at port 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil{
		log.Fatal(err)
	}
}