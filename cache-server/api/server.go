package server

import(
	"net/http"
)

func handleGetByUid(rw http.ResponseWriter, r *http.Response){
	if r.Request.Method != http.MethodGet{
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("json"))
}