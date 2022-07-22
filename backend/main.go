package main

import (
	// "embed"
	"html/template"
	// "io/fs"
	"log"
	"net/http"
	"path"

	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()

	var dataAkun akun.DataAkun
	dataAkun = akun.NewDataAkun()

	router.HandleFunc("/akun/login", dataAkun.LoginHandler)
	router.HandleFunc("/akun/register", dataAkun.RegisterHandler)
	router.HandleFunc("/home", dataAkun.HomeHandler)
	router.HandleFunc("/", LoginHandler)

	err := http.ListenAndServe(":8080", router)
	if err == nil{
		log.Fatal(err)
	}
	log.Println("SERVER is running at port 8080")
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	tmplt, err := template.ParseFiles(path.Join("views", "home.html"))
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

func LoginHandler(w http.ResponseWriter, r *http.Request){
	tmplt, err := template.ParseFiles("C:/Users/USER/Desktop/Pemrograman/Golang/HelloCoffee-with-golang/backend/views/login.html")
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
