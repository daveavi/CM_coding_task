package main 

import (
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/r3labs/sse/v2" 
	"os"
	"log"
)


func handleSSE(){

	type TestData struct {
		StudentId string `json:"studentId"`
		Exam int `json:"exam"`
		Score float64 `json:"score"`
	}

	client := sse.NewClient("http://live-test-scores.herokuapp.com/scores")

	client.Subscribe("messages", func(msg *sse.Event){
		var testData TestData 
		json.Unmarshal(msg.Data, &testData)		
		
		examString := strconv.Itoa(testData.Exam)
		scoreString := fmt.Sprintf("%f", testData.Score)
		writeDataToFile(testData.StudentId, examString, scoreString)


		mutexStudent.Lock()
		insertIntoStudents(testData.StudentId, testData.Score)
		mutexStudent.Unlock()

		mutexExam.Lock()
		insertIntoExams(examString, testData.Score)
		mutexExam.Unlock()

		
		
	})
}


func writeDataToFile(studentId string, examString string, scoreString string){
	logData := []byte("Student Id is: " + studentId +"\n" + "Exam Number is: " +examString +"\n" + "Score is: " + scoreString + "\n\n")

	f, err := os.OpenFile("examStudentLog.txt",  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		log.Fatal(err)
	}

	if _, err := f.Write(logData); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}


func insertIntoStudents(id string, score float64){
	studentToMarks[id] = append(studentToMarks[id], score)
}

func insertIntoExams(number string, score float64){	
	examToMarks[number] = append(examToMarks[number], score)
}
