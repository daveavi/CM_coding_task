ChannelMeter Coding Assessment by Avi Dave

Running Application: 
    - I have provided two bash scripts in the root directory to start the server
        1. startWithSSE.sh : This script kickstarts the server along with calling the Heroku service to send Server-Sent Events in a goroutine
        2. startNoSSE.sh: This script kickstarts the server without calling the Heroku service. 




Testing:
    - For testing, I provided a series of integration tests and a python script to test the performance of the get requests 

    -It is a bit difficult to test for functionality given that the stream is constantly flowing through our API. So to test for endpoints /exams/{number} and /students/{id}, I have made my SSE client create and write to a file under the app folder called, examStudentLog.txt, that will be logging data coming in via the JSON payload. That way it will be possible for you to test specific IDs. I used the text file examStudentLog.txt to manually update the curl calls in my bash scripts so I can test for 200 status responses.


    Integration Tests:
        1. testExamMarksList.sh: This test is to ensure that the marks of a specified exam paper get listed along with the average across all students, and see if it gets updated throughout time. 
        
        2. testExamNotExist.sh: This test is to validate if we get a 404 if we try to request marks for an exam that doesn't exist.

        3. testExamList.sh: This test is to ensure if we output the list of exams that have come in from the Heroku service, and does it get bigger throughout time
        
        3. testStudentList.sh: Same as testExamList.sh, but for students.
        
        4. testStudentMarksList.sh: Same test as testExamMarksList, but this is to ensure that exam marks for the specific student are listed as well as the average of their marks changing over time. 

        5. testStudentNotExist.sh: This test is to validate if we get a 404 if we try to request marks for a student that doesn't exist. 

        Tests testExamMarksList.sh,testExamList,testStudentMarksList,testStudentList, I ran startWithSSE.sh in a separate terminal or bash window, unfortunately, there was no command I could find that could replicate this action using macOS terminal. For tests testExamNotExist and testStudentNotExist, I ran startNoSSE.sh in another terminal window so no data will be populated in my local memory space, to test if I get a 404 accordingly.  

    Performance Tests: 
        1. readTest.py: This test is used to test the performance of doing concurrent get requests to my API. To run this test, run the command python3 readTest.py. To change the number of requests, you can increase the requests variable to a higher number, right now it is at 500. Also, python3 is required to run this script, if you are using a mac, you can simply brew install python, and you should be able to use python3 in your terminal. 

        Currently, the bottleneck for several concurrent requests is around 500, after 500 requests, the test starts crashing. I was having trouble figuring this out because I was thinking my bottleneck could be my HTTP request handler, but I was confused because go handles HTTP requests concurrently.  
    


Before running startWithSSE.sh or startNoSSE.SH:
    - run this command on terminal to give them executable permissions: chmod +x {"scriptName"}. After that simply run the script by typing in the name of the script. 
    