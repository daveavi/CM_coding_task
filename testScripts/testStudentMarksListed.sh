sleep 2 

echo First Student List:

echo 
curl http://localhost:8081/students/Vernie77 

sleep 5 

echo Second Student List:

echo 

curl http://localhost:8081/students/Vernie77

port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}
