package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:embed webstatic/*
var htmlData embed.FS

func main(){
	router := mux.NewRouter()

	serverRoot, err := fs.Sub(htmlData, "webstatic")
	if err != nil{
		log.Fatal(err)
	}

	var dataAkun akun.DataAkun
	dataAkun = akun.NewDataAkun()



	router.PathPrefix("/").Handler(http.FileServer(http.FS(serverRoot)))

	headersOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedOrigins([]string{"GET","HEAD","POST","PUT","OPTIONS"})

	log.Println("SERVER is running at port 8080")
	err = http.ListenAndServe(":8080", handlers.CORS(headersOk, methodsOk)(router))
	if err != nil{
		log.Fatal(err)
	}
}