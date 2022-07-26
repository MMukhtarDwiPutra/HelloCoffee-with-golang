package akun

import "net/http"

type DataAkun interface{
	// GetAkun(w http.ResponseWriter, r *http.Request)
	RegisterNewAccount(w http.ResponseWriter, r *http.Request)
	HomeHandler(w http.ResponseWriter, r *http.Request)
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
	SettingHandler(w http.ResponseWriter, r *http.Request)
	DeleteAkun (w http.ResponseWriter, r *http.Request)
	EditAkunHandler(w http.ResponseWriter, r *http.Request)
	EditAkun(w http.ResponseWriter, r *http.Request)
	EditPasswordHandler(w http.ResponseWriter, r *http.Request)
	EditPassword(w http.ResponseWriter, r *http.Request)
}