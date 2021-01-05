sleep 2 

echo First Exam List:

echo 
curl http://localhost:8081/exams 

sleep 2

echo Second Exam List:

echo 
curl http://localhost:8081/exams/500 

port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}