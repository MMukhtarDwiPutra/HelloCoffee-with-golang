package main

import (
	// "embed"
	"html/template"
	// "io/fs"
	"log"
	"net/http"

	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	// "github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/model"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"
)

func main(){
	router := mux.NewRouter()

	var dataAkun akun.DataAkun
	dataAkun = akun.NewDataAkun()

	router.HandleFunc("/akun/login", dataAkun.LoginHandler)
	router.HandleFunc("/akun/register", dataAkun.RegisterHandler)
	router.HandleFunc("/home", dataAkun.HomeHandler)
	router.HandleFunc("/akun/registration", dataAkun.RegisterNewAccount)
	router.HandleFunc("/", LoginHandler)
	router.HandleFunc("/akun/pengaturan/", dataAkun.SettingHandler)
	router.HandleFunc("/logout", dataAkun.LogoutHandler)
	router.HandleFunc("/akun/pengaturan/deleteAkun/", dataAkun.DeleteAkun)
	router.HandleFunc("/akun/pengaturan/edit/", dataAkun.EditAkunHandler)
	router.HandleFunc("/akun/pengaturan/edit/process/",dataAkun.EditAkun)
	router.HandleFunc("/akun/pengaturan/edit/password/", dataAkun.EditPasswordHandler)
	router.HandleFunc("/akun/pengaturan/edit/password/process/",dataAkun.EditPassword)
	router.HandleFunc("/toko/", dataAkun.DetailTokoHandler)
	router.HandleFunc("/menu", dataAkun.MenuHandler)
	router.HandleFunc("/menu/detail/", dataAkun.DetailMenu)
	router.HandleFunc("/komentar/tambahKomentar/", dataAkun.TambahKomentar)
	router.HandleFunc("/komentar/hapusKomentar/", dataAkun.HapusKomentar)
	router.HandleFunc("/keranjang/", dataAkun.KeranjangHandler)
	router.HandleFunc("/keranjang/tambahKeranjang/", dataAkun.TambahKeranjang)
	router.HandleFunc("/keranjang/hapusSemua/", dataAkun.HapusKeranjang)
	router.HandleFunc("/keranjang/checkout/", dataAkun.CheckoutHandler)
	router.HandleFunc("/checkout/process/", dataAkun.CheckoutProcess)
	router.HandleFunc("/home/toko", dataAkun.HomeTokoHandler)
	router.HandleFunc("/menu/edit/", dataAkun.EditMenuHandler)
	router.HandleFunc("/menu/hapus/",dataAkun.DeleteMenu)
	router.HandleFunc("/menu/tambahMenu", dataAkun.TambahMenuHandler)
	router.HandleFunc("/menu/tambahMenu/process", dataAkun.TambahMenuProcess)
	router.HandleFunc("/menu/edit/process/", dataAkun.EditMenuProcess)
	router.HandleFunc("/transaksi/", dataAkun.TransaksiHandler)
	router.HandleFunc("/transaksi/process/", dataAkun.ProcessTransaksi)

	path, _ := os.Getwd()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(path+`\backend\assets`))))
    
	log.Println("SERVER is running at port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil{
		log.Fatal(err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	path, _ := os.Getwd()
	tmplt, err := template.ParseFiles(path+`\backend\views\login.html`)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	err = tmplt.Execute(w, nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}	
}