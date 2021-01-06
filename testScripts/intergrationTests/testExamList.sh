echo First List of Exams:
curl http://localhost:8081/exams 
echo 
sleep 5


echo Second List of Exams:
curl http://localhost:8081/exams


port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}