package main

import (
	"log"
	"net/http"
)

var (
	registrator Registrator
)

func registerHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	log.Println("Register called")
	registrator.Register()

}

func unregisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("unregister called")
	registrator.Unregister()
}

func main() {

	registrator = &ConsulSvcRegistrator{}

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/unregister", unregisterHandler)
	http.ListenAndServe(":16000", nil)

}
