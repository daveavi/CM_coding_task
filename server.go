package main 

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/r3labs/sse/v2" 
)





//need to create our event Source handler as well
//need to also get clear with go programming to make this easier, as well as, really invest time to think
func allStudents(w http.ResponseWriter, r *http.Request){
	fmt.Println("About to extract all students")
}

func studentMarks(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	
	fmt.Println("About to extract exam marks for student " + id)
}

func allExams(w http.ResponseWriter, r *http.Request){

	fmt.Println("About to extract all exams")
}
func examMarks(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	number := vars["number"]

	fmt.Println("About to extract marks for exam number "+ number)
}

func handleRequests(){

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/students", allStudents)
	
	myRouter.HandleFunc("/students/{id}", studentMarks)
	myRouter.HandleFunc("/exams", allExams)
	myRouter.HandleFunc("/exams/{number}", examMarks)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	type TestData struct {
		StudentId string `json:"studentId"`
		Exam int `json:"exam"`
		Score float64 `json:"score"`
	}

	client := sse.NewClient("http://live-test-scores.herokuapp.com/scores")

	client.Subscribe("messages", func(msg *sse.Event){
		var testData TestData 
		json.Unmarshal(msg.Data, &testData)		
		fmt.Println(testData)
	})



	handleRequests()
}