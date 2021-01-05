ChannelMeter Coding Assessment by Avi Dave

Running Application: 
    - I have provided two bash scripts in the root directory to start the server
        1. startWithSSE.sh : This script kickstarts the server along with waiting for the heroku service to send Server Sent Events in a goroutine
        2. startNoSSE.sh: This script kickstarts the server without calling the heroku service. 




Testing:
    - For testing I provided a series of integration tests and a python script to test the performance of the get requests 

    -It is a bit difficult to test for functionality given that the stream is constantly flowing through our API. So to test for endpoints /exams/{number} and /students/{id},I have made my SSE client call back create and write to a file under the app folder called, examStudentLog.txt, that will be logging data coming in via the JSON payload. That way it will be possible for you test specific data. I used the text file examStudentLog.txt to manually update the curl calls in my bash scripts so I can test for 200 status responses. I know this isn't the best practice, because it would be better to find a way to automate it somehow, but with the stream constantly flowing through, it was hard for me to test for an example that never changes.


    Integration Tests:
        1. testExamMarksList.sh: This test is to ensure that the marks of a specified exam paper get listed a long with the average across all students ,and seeing if it get's updated throughout time. 
        
        2. testExamNotExist.sh: This test is to validate if we get a 404 if we try to request marks for an exam that doesn't exist.

        3. testExamList.sh: This test is to ensure if we output the list of exams that has come in from the heroku service, and does it get bigger throughout time
        
        3. testStudentList.sh: Same as testExamList.sh, but for students.
        
        4. testStudentMarksList.sh: Almost the same test as testExamMarksList, but this is to ensure that exam marks for the specific student is listed as well as the average of their marks changing over time. 

        5. testStudentNotExist.sh: This test is to validate if we get a 404 if we try to request marks for a student that doesn't exist. 

        Tests testExamMarksList.sh,testExamList,testStudentMarksList,testStudentList require you to run startWithSSE.sh in seperate terminal or bash window, unfortunatley there was no command I could find that could replicate this action using MacOS terminal. For tests testExamNotExist and testStudentNotExist, run startNoSSE.sh in another terminal so no data will be populated in my local memory space.  

    Performance Tests: 
        1. readTest.py: This test is used to test the performance of doing concurrent get requests to my Api. To run this test, run the command python3 readTest.py. To change the number of requests, you can increase the requests variable to a higher number, right now it is at 100. Also python3 is required to run this script, if you are using a mac, you can simply brew install python, and you should be able to use python3 in terminal. 

        Currently the bottleneck for number of conccurent requests is at 500, so it is likely that after 500 requests, it stops working.
    





Before running any bashscripts, make sure to run this command on terminal to give them executable permissions: chmod -x {"scriptName"}. After that simply run the script by typing in the name of the script.