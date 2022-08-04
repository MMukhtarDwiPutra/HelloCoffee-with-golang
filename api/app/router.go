package app

import(
	"github.com/gorilla/mux"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/controller"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
)

func AddRouterAPI(router *mux.Router, menuService menu.Service) *mux.Router{
	menuController := controller.NewMenuController(menuService)

	router.HandleFunc("/api/menu", menuController.FindAllMenu).Methods("GET")
	router.HandleFunc("/api/menu/{id}", menuController.FindMenu).Methods("GET")
	router.HandleFunc("/api/menu", menuController.CreateMenu).Methods("POST")
	router.HandleFunc("/api/menu/{id}", menuController.UpdateMenu).Methods("PUT")
	router.HandleFunc("/api/menu/{id}", menuController.DeleteMenu).Methods("DELETE")

	return router
}