package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)





//need to create our event Source handler as well
//need to also get clear with go programming to make this easier, as well as, really invest time to think
func allStudents(w http.ResponseWriter, r *http.Request){
	fmt.Println("About to extract all students")
	//try to get the students table and lock it
	//go through key and value and fprintf each student 
	//release lock

	if len(studentToMarks) == 0{
		fmt.Fprintf(w, "No student's marks have been recorded"+"\n")
	}else{
		for key,_ := range studentToMarks{
			fmt.Fprintf(w, key+"\n")
		}
	}


}
func studentMarks(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	
	fmt.Println("About to extract exam marks for student " + id)

	//go through the list that is mapped to this student id
	//add all the scores from this list and take average 
	//then fprintf the average 

	//get lock first
	if marks, ok:=studentToMarks[id]; ok{
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
		fmt.Fprintf(w, "Student " + id+" does not exist\n")
	}



	fmt.Println("About to extract exam marks for student " + id)
}

func allExams(w http.ResponseWriter, r *http.Request){
	//same as all allstudents 
	fmt.Println("About to extract all exams")
	//get Locks
	if len(examToMarks) == 0{
		fmt.Fprintf(w,"No exam marks have been recorded yet" +"\n")
	}else{
		for key,_ := range examToMarks{
			fmt.Fprintf(w, key+"\n")
		}
	}
}


func examMarks(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	number := vars["number"]


	//get lock first 
	if marks, ok := examToMarks[number]; ok{
		average := 0.0
		
		for _,num := range marks{
			s := fmt.Sprintf("%f", num)
			fmt.Fprintf(w, s +"\n")
			average += num
		}

		length := len(examToMarks[number])
		average = average / float64(length)

		strAverage := fmt.Sprintf("%f", average)
		fmt.Fprintf(w, "For exam: " + number+", the average turned out to be"+ strAverage+ "\n")
	}else{
		fmt.Fprintf(w, "Exam " + number+" does not exist\n")
	}

	
	//same as all students
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
	studentToMarks = make(map[string][]float64)
	examToMarks = make(map[string][]float64)

	studentToMarks["Kellie64"] = append(studentToMarks["Kellie64"] , 0.743508)
	studentToMarks["Kellie64"] = append(studentToMarks["Kellie64"] , 0.764371)
	
	go func(){
		handleSSE()
	}() 


	handleRequests()

}