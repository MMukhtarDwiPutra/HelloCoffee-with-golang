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
	router.HandleFunc("/", LoginHandler)

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