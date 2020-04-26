package main

import (
	"./app"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Anasayfaya Hoşgeldiniz.")
	fmt.Println("Anasayfa Servisi Çağırıldı")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Anasayfa).Methods("GET")
	router.HandleFunc("/anasayfa", Anasayfa).Methods("GET")

	router.Use(app.JwtAuthentication)
	http.ListenAndServe(":5555", router)
}

func main() {
	fmt.Println("Go Dili İle İlk Service Uygulaması Başlatılıyor...")
	handleRequest()
}
