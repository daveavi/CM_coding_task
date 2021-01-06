echo List of Students:
curl http://localhost:8081/students 
echo 


echo "AviDave's" Exam Marks:
curl http://localhost:8081/students/AviDave 

port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}