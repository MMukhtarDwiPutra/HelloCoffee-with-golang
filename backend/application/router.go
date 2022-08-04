package application

import(	
	"github.com/MMukhtarDwiPutra/HelloCoffee-with-golang/backend/handler"
	"github.com/gorilla/mux"
	"os"
	"net/http"
)


func NewRouter(akunHandler *handler.AkunHandler, tokoHandler *handler.TokoHandler, menuHandler *handler.MenuHandler, keranjangHandler *handler.KeranjangHandler) *mux.Router{
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

	return router
}