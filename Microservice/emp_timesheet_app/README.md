__Steps to build and run the Docker container__

_For User Login Service_

cd emp-login-service
docker build -t emp-login-service .
docker run -p 8080:8080 emp-login-service

_For Timesheet Service_

cd timesheet-service
docker build -t timesheet-service .
docker run -p 8081:8081 timesheet-service


__Steps for Jenkins LTS__

Install the latest LTS version: _brew install jenkins-lts_
Start the Jenkins service: _brew services start jenkins-lts_
Restart the Jenkins service: _brew services restart jenkins-lts_
Update the Jenkins version: _brew upgrade jenkins-lts_


After starting the Jenkins service, browse to http://localhost:8080 and follow the instructions to complete the installation.

copy the password in the prompt from below location.
/Users/anusharesu/.jenkins/secrets/initialAdminPassword

