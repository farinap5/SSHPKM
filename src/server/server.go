package server

import (
	"context"
	"fmt"
	"github.com/farinap5/SSHPKM/src/db"
	"net/http"
)

var server *http.Server
var serverMux *http.ServeMux

func ServerPrepare() {
	serverMux = http.NewServeMux()
	serverMux.HandleFunc("/", Home)

	Addr := db.DBGetConfig("Address") + ":" + db.DBGetConfig("Port")

	server = &http.Server{
		Addr:    Addr,
		Handler: serverMux,
	}
}

func ServerStart() {
	fmt.Println("Listener requested for: " + server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}

}

func ServerListen() bool {
	ServerPrepare()
	go ServerStart()
	fmt.Println("Starting server: " + server.Addr)
	return true
}

func ServerStop() bool {
	fmt.Println("Shutdown requested.")
	if server == nil {
		fmt.Println("Server is not started.")
		return false
	}
	err := server.Shutdown(context.Background())
	fmt.Println("Stopping server: " + server.Addr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func ServerRestart() {
	fmt.Println("Restarting...")
	ServerStop()
	ServerListen()
}
