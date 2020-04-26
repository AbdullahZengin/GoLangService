package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Anasayfaya Hoşgeldiniz.")
	fmt.Println("Anasayfa Servisi Çağırıldı")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Anasayfa).Methods("GET")
	http.ListenAndServe(":5555", router)
}

func main() {
	fmt.Println("Go Dili İle İlk Service Uygulaması Başlatılıyor...")
	handleRequest()
}
