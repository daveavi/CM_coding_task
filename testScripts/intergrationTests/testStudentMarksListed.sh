echo "Reece_Wisoky6's" Exam Marks:
curl http://localhost:8081/students/Reece_Wisoky6 
echo

sleep 5 

echo "Reece_Wisoky6's" Exam Marks updated:
curl http://localhost:8081/students/Reece_Wisoky6

port=8081
result="`lsof -Fp -n -i :$port | grep p`"
kill -9 ${result##p}
