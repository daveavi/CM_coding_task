#note, the exam number may need to change for the test
echo Marks "for" Exam 8489:
curl http://localhost:8081/exams/8489 
echo
sleep 5 


echo Exam 8489 marks updated:
curl http://localhost:8081/exams/8489

port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}
