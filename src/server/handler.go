package server

import (
	"github.com/farinap5/SSHPKM/src/db"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	Host := r.Header.Get("SSH-Host")
	User := r.Header.Get("SSH-User")

	if (Host == "" && User == "") || Host == "" || User == "" {
		w.Write([]byte("(null)"))
		return
	}

	v, _ := db.DBVerifyAccess(User, Host)
	if !v {
		w.Write([]byte("(null)"))
		return
	}

	pk := db.DBGetUserSSHKey(User) + "\n"
	w.Write([]byte(pk))
}
