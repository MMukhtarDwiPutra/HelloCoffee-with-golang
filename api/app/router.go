package app

import(
	"github.com/gorilla/mux"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/controller"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/exception"
)

func AddRouterAPI(router *mux.Router, menuService menu.Service, userService akun.Service) *mux.Router{
	menuController := controller.NewMenuController(menuService)
	userController := controller.NewUserController(userService)

	router.HandleFunc("/api/menu", menuController.FindAllMenu).Methods("GET")
	router.HandleFunc("/api/menu/{id}", menuController.FindMenu).Methods("GET")
	router.HandleFunc("/api/menu", menuController.CreateMenu).Methods("POST")
	router.HandleFunc("/api/menu/{id}", menuController.UpdateMenu).Methods("PUT")
	router.HandleFunc("/api/menu/{id}", menuController.DeleteMenu).Methods("DELETE")

	router.HandleFunc("/api/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", userController.FindAllUser).Methods("GET")
	router.HandleFunc("/api/user/{id}", userController.FindById).Methods("GET")
	router.HandleFunc("/api/user/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", userController.DeleteUser).Methods("DELETE")

	router.Use(exception.ErrorHandler)
	return router
}