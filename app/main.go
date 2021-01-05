package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
)


func allStudents(w http.ResponseWriter, r *http.Request){
	fmt.Println("About to extract all students")
	mutexStudent.Lock()
	if len(studentToMarks) == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - NotFound\n"))
		fmt.Fprintf(w, "No student's marks have been recorded"+"\n")
	}else{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Okay\n"))

		for key,_ := range studentToMarks{
			fmt.Fprintf(w, key+"\n")
		}

	}
	mutexStudent.Unlock()
}

func studentMarks(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	
	fmt.Println("About to extract exam marks for student " + id)

	mutexStudent.Lock()
	if marks, ok:=studentToMarks[id]; ok{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Okay\n"))
		average := 0.0 
		
		for _,num := range marks{
			s := fmt.Sprintf("%f", num)
			fmt.Fprintf(w, s +"\n")
			average += num
		}

		length := len(studentToMarks[id])
		average = average / float64(length)
		strAverage := fmt.Sprintf("%f", average)

		fmt.Fprintf(w, "Student " + id+" exam marks average to "+ strAverage+ "\n")
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - NotFound\n"))
		w.Write([]byte("Student " + id+" does not exist\n"))
	}
	mutexStudent.Unlock()

}


func allExams(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("About to extract all exams")
	mutexExam.Lock()
	if len(examToMarks) == 0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - NotFound\n"))
		w.Write([]byte("No exam marks have been recorded yet" +"\n"))
	}else{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Okay\n"))

		for key,_ := range examToMarks{
			w.Write([]byte(key+"\n"))
		}
	}
	mutexExam.Unlock()
}


func examMarks(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	number := vars["number"]

	fmt.Println("About to extract marks for exam number "+ number)

	mutexExam.Lock()
	if marks, ok := examToMarks[number]; ok{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Okay\n"))
		average := 0.0
		
		for _,num := range marks{
			s := fmt.Sprintf("%f", num)
			w.Write([]byte(s +"\n"))
			average += num
		}

		length := len(examToMarks[number])
		average = average / float64(length)
		strAverage := fmt.Sprintf("%f", average)

		w.Write([]byte("For exam: " + number+", the average turned out to be "+ strAverage+ "\n"))
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - NotFound\n"))
		w.Write([]byte("Exam " + number+" does not exist\n"))
	}
	mutexExam.Unlock()

	
}






func handleRequests(){
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/students", allStudents)
	myRouter.HandleFunc("/students/{id}", studentMarks)
	myRouter.HandleFunc("/exams", allExams)
	myRouter.HandleFunc("/exams/{number}", examMarks)

	fmt.Printf("Connecting to port 8081\n")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main(){
	

	studentToMarks = make(map[string][]float64)
	examToMarks = make(map[string][]float64)


	toGetAllArgs := os.Args[1:] 

	if len(toGetAllArgs) == 1 && toGetAllArgs[0] == "startSSE" {
		go func(){
			fmt.Printf("Starting goroutine, for Server Sent Events\n")
			handleSSE()
		}() 
	}else if len(toGetAllArgs) > 1{
		fmt.Printf("You only need one argument, to run this server\n")
		fmt.Printf("If you want to start Server Sent Events, only add startSSE\n")
		return
	}

	handleRequests()

}