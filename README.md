# Simple Maths Sum API in go

# Play the quiz with curl:

## Entry to the quiz
```
curl -XGET 'http://localhost:8080/'
```
## Start a new quiz
```
curl -XPOST 'http://localhost:8080/new'
```
## Submit an answer to a quiz question
```
curl -XPOST 'http://localhost:8080/quiz/1'
```
## Get the current status of your quiz
```
curl -XGET 'http://localhost:8080/quiz/1'
```
## Overview of the state of all quizes
```
curl -XGET 'http://localhost:8080/quizes/'
```
