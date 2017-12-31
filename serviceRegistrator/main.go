package main

import (
	"log"
	"net/http"
	"os"
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
	log.Println("unregister called1")
	registrator.Unregister()
}
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("ping called")
	registrator.Unregister()
}

func main() {

	registrator = &ConsulSvcRegistrator{}

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/unregister", unregisterHandler)
	http.HandleFunc("/ping", pingHandler)
	err := registrator.Register()
	if err != nil {
		log.Fatal("Could not register service")
		panic(err)
	}
	host, _ := os.Hostname()
	http.ListenAndServe(host+":9090", nil)

}
