# Simple Maths Sum API in go

# Run the quiz API:
```
git clone git@github.com:pidah/simple-sums-api.git
cd simple-sums-api
go run *.go
```
# Play the quiz with curl:

## Entry to the quiz
```
$ curl -XGET 'http://localhost:8080/'
Welcome to the Math quiz game!
```
## Start a new quiz
```
$ curl -XGET 'http://localhost:8080/new'
{"id":1,"first_number":9,"second_number":7,"score":0,"correct":false}
$ curl -XGET 'http://localhost:8080/new'
{"id":2,"first_number":7,"second_number":3,"score":0,"correct":false}
```
## Submit an answer to a quiz question
```
$ curl -H "Content-Type: application/json" -X POST -d '{"data":16}' http://localhost:8080/quiz/1
{"id":1,"first_number":9,"second_number":7,"score":1,"correct":true}
```
## Get the current status of your quiz
```
curl -XGET 'http://localhost:8080/quiz/1'
{"id":1,"first_number":5,"second_number":7,"score":1,"correct":false}
```
## Overview of the state of all quizes in-play
```
curl -XGET 'http://localhost:8080/quizes/'
[{"id":1,"first_number":5,"second_number":7,"score":1,"correct":false},{"id":2,"first_number":5,"second_number":2,"score":0,"correct":false}]
```
