package akun

import "net/http"

type DataAkun interface{
	// GetAkun(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
	SettingHandler(w http.ResponseWriter, r *http.Request)
	DeleteAkun (w http.ResponseWriter, r *http.Request)
	EditAkunHandler(w http.ResponseWriter, r *http.Request)
	EditAkun(w http.ResponseWriter, r *http.Request)
	EditPasswordHandler(w http.ResponseWriter, r *http.Request)
	EditPassword(w http.ResponseWriter, r *http.Request)
	TambahKomentar(w http.ResponseWriter, r *http.Request)
	HapusKomentar(w http.ResponseWriter, r *http.Request)
	KeranjangHandler(w http.ResponseWriter, r *http.Request)
	TambahKeranjang(w http.ResponseWriter, r *http.Request)
	HapusKeranjang(w http.ResponseWriter, r *http.Request)
	CheckoutHandler (w http.ResponseWriter, r *http.Request)
	CheckoutProcess (w http.ResponseWriter, r *http.Request)
	HomeTokoHandler (w http.ResponseWriter, r *http.Request)	
	TransaksiHandler (w http.ResponseWriter, r *http.Request)
	ProcessTransaksi(w http.ResponseWriter, r *http.Request)
	CheckoutNowHandler(w http.ResponseWriter, r *http.Request)
}