echo List of Exams:
curl http://localhost:8081/exams 
echo

sleep 2

echo Exam Marks "for" 1avidave:
curl http://localhost:8081/exams/1avidave

port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}