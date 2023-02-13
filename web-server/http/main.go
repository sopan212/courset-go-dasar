package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
}

type Articles []Article
type Employee struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
type Employes []Employee

func main() {
	fmt.Println("http")
	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/article/submit", postArticles)
	http.HandleFunc("/employe", getEmployes)
	http.HandleFunc("/employe/submit", postEmployes)
	log.Printf("service run in localhost:3000")
	http.ListenAndServe(":3000", nil)
}

func getEmployes(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET: /employe")
	employe := Employes{
		Employee{Name: "sopan", Phone: "08782031313", Address: "jakarta"},
	}
	json.NewEncoder(w).Encode(employe)
}

func postEmployes(w http.ResponseWriter, r *http.Request) {

	log.Printf("POST: /employes/submit")
	if r.Method == "POST" {
		var employes Employee
		err := json.NewDecoder(r.Body).Decode(&employes)
		if err != nil {
			log.Printf("Bad Request", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if employes.Name == "" || employes.Address == "" || employes.Phone == "" {
			log.Printf("nama alamat dan phone tidak boleh kosong")
			http.Error(w, "nama alamat dan phone tidak boleh kosong", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(employes)
	} else {
		log.Printf("invalid method")
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
		return
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halaman Home"))
}
func getArticles(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET: /articles")
	articles := Articles{
		Article{Judul: "test judul",
			Penulis: "test penulis",
		},
	}
	json.NewEncoder(w).Encode(articles)
}

func postArticles(w http.ResponseWriter, r *http.Request) {

	log.Printf("POST: /articles")
	if r.Method == "POST" {
		var article Article
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Printf("Bad Request", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if article.Judul == "" || article.Penulis == "" {
			log.Printf("judul atau penulis tidak boleh kosong")
			http.Error(w, "Bad Request,judul atau penulis tidak boleh kosong", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(article)
	} else {
		log.Printf("invalid method")
		http.Error(w, "invalid Method", http.StatusMethodNotAllowed)
		return
	}
}
