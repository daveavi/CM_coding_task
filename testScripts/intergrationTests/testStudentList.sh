echo First List of Students:
curl http://localhost:8081/students 
echo

sleep 2 


echo Second List of Students:
curl http://localhost:8081/students


port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}