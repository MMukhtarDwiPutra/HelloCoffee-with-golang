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

	// http.Handle("/css/", http.StripPrefix("C:/Users/USER/Desktop/Pemrograman/Golang/HelloCoffee-with-golang/backend/views/css/", http.FileServer(http.Dir("C:/Users/USER/Desktop/Pemrograman/Golang/HelloCoffee-with-golang/backend/views/css"))))

	log.Println("SERVER is running at port 8080")
	err := http.ListenAndServe(":8080", router)
	if err == nil{
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