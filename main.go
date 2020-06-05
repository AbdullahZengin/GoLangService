package main

import (
	"./app"
	"./models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var Articles []models.Article

func Anasayfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Anasayfaya Hoşgeldiniz.")
	fmt.Println("Anasayfa Servisi Çağırıldı")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Makaleler endpoint'i çağırıldı.")
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article models.Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Anasayfa).Methods("GET")
	router.HandleFunc("/anasayfa", Anasayfa).Methods("GET")
	router.HandleFunc("/TumMakaleleriGetir", returnAllArticles).Methods("GET")
	router.HandleFunc("/YeniMakale", createNewArticle).Methods("POST")
	router.Use(app.LoginMiddleWare)
	http.ListenAndServe(":5555", router)

}

func main() {

	Articles = []models.Article{
		models.Article{Id: "1", Title: "Golang Nedir ?", Desc: "Golang hakkında kısa bir bilgi.", Content: "Go veya Golang, 2007 yılında Google’ın geliştirmeye başladığı açık kaynak programlama dilidir. Kullanım alanı olarak daha çok sistem programlama için kullanılmaktadır. 2009 yılı Kasım ayında çıkmıştır.  Go derleyicisi “gc” açık kaynak yazılım olarak; Windows, Linux, OS X, BSD ve Unix versiyonları geliştirilmiştir. 2015 yılından beri de akıllı telefonlar için geliştirilmeye başlanmıştır.\n\nGo, Google mühendisleri tarafından deney olarak ortaya çıkarılmıştır. Diğer dillerin bilinen eleştirilerini çözecek biçimde tasarlanmıştır. Aynı zamanda olumlu özelliklerini de koruyacak şekilde geliştirilmiştir. Go programlama dili; üretken ve okunabilir olması, ağ ve çoklu işlemleri desteklemesi, statik yazılmış ve büyük sistemlere ölçeklenebilir olması özelliklerini taşıyordu. Golang 2007 yılında ilk adımı atılan, 2009 yılında dile getirilen, 2012 ortalarında Go 1.0’a ulaşan bir programlama dilidir. Google tarafından desteklenen, basit ve sunduğu performansla öne çıkan, oldukça genç, açık kaynak bir programlama dilidir."},
		models.Article{Id: "2", Title: "Rest Service Nedir ?", Desc: "Rest service hakkında tanımlayıcı bilgi.", Content: "REST, servis yönelimli mimari üzerine oluşturulan yazılımlarda kullanılan bir veri transfer yöntemidir. HTTP üzerinde çalışır ve diğer alternatiflere göre daha basittir, minimum içerikle veri alıp gönderdiği için de daha hızlıdır. İstemci ve sunucu arasında XML veya JSON verilerini taşıyarak uygulamaların haberleşmesini sağlar. REST standartlarına uygun yazılan web servislerine RESTful servisler diyoruz. "},
	}
	fmt.Println("Golang ile rest service uygulaması başlatılıyor...\nPort: 5555")
	handleRequest()
}
