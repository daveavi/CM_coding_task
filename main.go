package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}


type Articles []Article

var articles = [2]Article{
	Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func allArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	// json.NewEncoder(w).Encode(articles)
	for _, article := range articles{
      json.NewEncoder(w).Encode(article)
    }

}

func getArticle(w http.ResponseWriter, r *http.Request){	
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Println("Endpoint Hit: Getting article "+key)


	for _, article := range articles{
		if article.Id == key{
			json.NewEncoder(w).Encode(article)
		}
	}

	fmt.Fprintf(w, "didn't find the article\n")

}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage Endpoint Hit\n")
}



func handleRequests(){
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/articles/{id}", getArticle)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}