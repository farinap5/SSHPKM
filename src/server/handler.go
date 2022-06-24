package server

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	r.URL.Query()
	w.Write([]byte("testeee"))
}
