package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to the Math quiz game!\n")
	rw.Header().Set("Content-Type", "application/json")
}

func StartQuiz(w http.ResponseWriter, r *http.Request) {
	q := NewQuiz(sessionID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(q); err != nil {
		panic(err)
	}
}

func MyQuiz(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/quiz/")
	var quizID int
	var err error
	if quizID, err = strconv.Atoi(id); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
	}

	switch r.Method {

	case "POST":

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		var answer Answer

		if err := json.Unmarshal(body, &answer); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(answer); err != nil {
				panic(err)
			}
			return
		}

		quiz := FindQuiz(quizID)
		if quiz.Result == answer.Data {
			quiz.Correct = true
			quiz.Score = quiz.Score + 1
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(quiz); err != nil {
				panic(err)
			}
			UpdateQuiz(quiz)
			return
		}

	default:

		quiz := FindQuiz(quizID)
		if quiz.ID > 0 {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(quiz); err != nil {
				panic(err)
			}
			return
		}

		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(quiz); err != nil {
			panic(err)
		}
	}
}

func QuizIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(quizes); err != nil {
		panic(err)
	}
}
