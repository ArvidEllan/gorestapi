package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}

	fmt.Println("Endpoint Hit: All articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

//func main() {
//	handleRequests()
//}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

//package main

//import (
//"encoding/json"
//"fmt"
//"log"
//"net/http"
//)

//type Article struct {

//Title string 'json:"Title"'
//Desc string 'json:"desc"'
//Content string 'json:"content"'

//}
//var Articles []Article

//func allArticles(w http.ResponseWriter, r *http.Request) {
//articles := Articles {
//Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
//}

//	fmt.Println("Endpoint Hit: All articles Endpoint")
//	json.NewEncoder(w).Encode(articles)

//}

//func homePage(w http.ResponseWriter, r *http.Request) {
//fmt.Fprint(w, "homepage Endpoint Hit")

//}

//func handleRequests() {

//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/articles", allArticles)
//	log.Fatal(http.ListenAndServe(":8081", nil))

//}

//func main() {

//	handleRequests()
//}
