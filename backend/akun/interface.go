package akun

import "net/http"

type DataAkun interface{
	getAkun(w http.ResponseWriter, r *http.Request)
}