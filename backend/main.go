package main

import (
	"database/sql"
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
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"
)

func main(){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/hellocoffee")
	if err != nil{
		log.Fatal(err)
	}

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

	router := mux.NewRouter()

	router.HandleFunc("/akun/login", akunHandler.LoginProcess)
	router.HandleFunc("/akun/register", akunHandler.RegisterHandler)
	router.HandleFunc("/akun/registration", akunHandler.RegisterNewAccount)
	router.HandleFunc("/", akunHandler.LoginHandler)

	router.HandleFunc("/home", tokoHandler.HomeHandler)
	router.HandleFunc("/toko/", tokoHandler.DetailTokoHandler)
	router.HandleFunc("/home/toko", tokoHandler.HomeTokoHandler)

	router.HandleFunc("/menu", menuHandler.MenuHandler)
	router.HandleFunc("/menu/detail/", menuHandler.DetailMenu)
	router.HandleFunc("/menu/edit/", menuHandler.EditMenuHandler)
	router.HandleFunc("/menu/tambahMenu", menuHandler.TambahMenuHandler)
	router.HandleFunc("/menu/tambahMenu/process", menuHandler.TambahMenuProcess)
	router.HandleFunc("/menu/hapus/", menuHandler.DeleteMenu)
	router.HandleFunc("/menu/edit/process/", menuHandler.EditMenuProcess)

	router.HandleFunc("/akun/pengaturan/", akunHandler.SettingHandler)
	router.HandleFunc("/logout", akunHandler.LogoutHandler)
	router.HandleFunc("/akun/pengaturan/deleteAkun/", akunHandler.DeleteAkun)
	router.HandleFunc("/akun/pengaturan/edit/", akunHandler.EditAkunHandler)
	router.HandleFunc("/akun/pengaturan/edit/process/",akunHandler.EditAkun)
	router.HandleFunc("/akun/pengaturan/edit/password/", akunHandler.EditPasswordHandler)
	router.HandleFunc("/akun/pengaturan/edit/password/process/",akunHandler.EditPassword)

	router.HandleFunc("/komentar/tambahKomentar/", akunHandler.TambahKomentar)
	router.HandleFunc("/komentar/hapusKomentar/", akunHandler.HapusKomentar)

	router.HandleFunc("/transaksi/", akunHandler.TransaksiHandler)
	router.HandleFunc("/transaksi/process/", akunHandler.ProcessTransaksi)
	
	router.HandleFunc("/keranjang/", keranjangHandler.KeranjangHandler)
	router.HandleFunc("/keranjang/tambahKeranjang/", keranjangHandler.TambahKeranjang)
	router.HandleFunc("/keranjang/hapusSemua/", keranjangHandler.HapusKeranjang)
	router.HandleFunc("/keranjang/checkout/", keranjangHandler.CheckoutHandler)
	router.HandleFunc("/checkout/process/", keranjangHandler.CheckoutProcess)
	router.HandleFunc("/keranjang/checkOutNow/", keranjangHandler.CheckoutNowHandler)

	path, _ := os.Getwd()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(path+`\backend\assets`))))
	router.PathPrefix("/static/img/").Handler(http.StripPrefix("/static/img/", http.FileServer(http.Dir(path+`\backend\assets\img`))))
    
	log.Println("SERVER is running at port 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil{
		log.Fatal(err)
	}
}