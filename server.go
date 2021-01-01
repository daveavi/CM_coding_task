package main 

import (
	"fmt"
	"log"
	"net/http"
)
//need to create our event Source handler as well
//need to also get clear with go programming to make this easier, as well as, really invest time to think
func allStudents(w http.ResponseWriter, r *http.Request){
	
}

func studentMarks(w http.ResponseWriter, r *http.Request){
	
}

func allExams(w http.ResponseWriter, r *http.Request){

}
func examMarks(w http.ResponseWriter, r *http.Request){
	
}

func handleRequests(){
	http.HandleFunc("/students", allStudents)
	http.HandleFunc("/students/{id}", studentMarks)
	http.HandleFunc("/exams", allExams)
	http.HandleFunc("/exams/{number}", examMarks)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}