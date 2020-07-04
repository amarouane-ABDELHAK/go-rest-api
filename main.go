package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"description"`
	Content string `json:"content"`
	
}

type Articles []Article
var articles Articles
func allArticles( w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Endpoint hits all articles endpoint")

	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		log.Fatal("Couldn't encode the stuff", err)
	}
}


func homePage( w http.ResponseWriter, r *http.Request) {
	if _, err:= fmt.Fprintf(w, "Homepage Endpoint hit"); err != nil {
		log.Fatal("Can't write into the response writer", err)
	}
}

func addArticle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "I got hit by a post request")
	var p Article
	json.NewDecoder(r.Body).Decode(&p)
	articles = append(articles, Article{
		Title:   p.Title,
		Desc:    p.Desc,
		Content: p.Content,
	})

	fmt.Println(p)
	
}
func handleRequest()  {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", addArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", myRouter))
}



func main()  {
	handleRequest()
}
