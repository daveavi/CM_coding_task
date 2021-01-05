sleep 2 

echo First Student List:

curl http://localhost:8081/exams 
echo 
sleep 2 


echo Second Student List:

curl http://localhost:8081/exams/500
echo 

port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}