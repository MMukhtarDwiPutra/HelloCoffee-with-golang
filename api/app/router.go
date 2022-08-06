package app

import(
	"github.com/gorilla/mux"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/controller"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/menu"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/toko"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/akun"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/keranjang"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/transaksi"
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/api/exception"
)

func AddRouterAPI(router *mux.Router, menuService menu.Service, userService akun.Service, tokoService toko.Service, keranjangService keranjang.Service, transaksiService transaksi.Service) *mux.Router{
	menuController := controller.NewMenuController(menuService)
	userController := controller.NewUserController(userService)
	tokoController := controller.NewTokoController(tokoService)
	keranjangController := controller.NewKeranjangController(keranjangService)
	transaksiController := controller.NewTransaksiController(transaksiService)

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

	router.HandleFunc("/api/toko", tokoController.CreateToko).Methods("POST")
	router.HandleFunc("/api/toko", tokoController.FindAllToko).Methods("GET")
	router.HandleFunc("/api/toko/{id}", tokoController.FindById).Methods("GET")
	router.HandleFunc("/api/toko/{id}", tokoController.UpdateToko).Methods("PUT")
	router.HandleFunc("/api/toko/{id}", tokoController.DeleteToko).Methods("DELETE")

	router.HandleFunc("/api/keranjang", keranjangController.CreateKeranjang).Methods("POST")
	router.HandleFunc("/api/keranjang", keranjangController.FindAllKeranjang).Methods("GET")
	router.HandleFunc("/api/keranjang/{id}", keranjangController.FindById).Methods("GET")
	router.HandleFunc("/api/keranjang/{id}", keranjangController.UpdateKeranjang).Methods("PUT")
	router.HandleFunc("/api/keranjang/{id}", keranjangController.DeleteKeranjang).Methods("DELETE")

	// router.HandleFunc("/api/transaksi", transaksiController.CreateTransaksi).Methods("POST")
	router.HandleFunc("/api/transaksi", transaksiController.FindAllTransaksi).Methods("GET")
	router.HandleFunc("/api/transaksi/{id}", transaksiController.FindById).Methods("GET")
	// router.HandleFunc("/api/transaksi/{id}", transaksiController.UpdateTransaksi).Methods("PUT")
	// router.HandleFunc("/api/transaksi/{id}", transaksiController.DeleteTransaksi).Methods("DELETE")

	router.Use(exception.ErrorHandler)
	return router
}