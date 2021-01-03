package SSE 

import (
	"fmt"
	"encoding/json"
	"github.com/r3labs/sse/v2" 
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
		
		fmt.Println("Student Id is: " + testData.StudentId)
		fmt.Printf("Exam Number is: %d\n" , testData.Exam) 
		fmt.Printf("Score is: %f\n", testData.Score)
	})
}

func main(){
	handleSSE()
}
