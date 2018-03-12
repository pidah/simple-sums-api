package main

import (
	"fmt"
	"math/rand"
)

var sessionID int

type Quiz struct {
	ID           int  `json:"id"`
	FirstNumber  int  `json:"first_number"`
	SecondNumber int  `json:"second_number"`
	Result       int  `json:"-"`
	Score        int  `json:"score"`
	Correct      bool `json:"correct"`
}

type Quizes []Quiz

var quizes Quizes

func NewQuiz(ID int) (q Quiz) {
	sessionID++
	q.ID = sessionID
	q.FirstNumber = rand.Intn(10)
	q.SecondNumber = rand.Intn(10)
	q.Score = 0
	q.CalcResult()
	q.Correct = false
	quizes = append(quizes, q)
	return
}

func (q *Quiz) CalcResult() {
	q.Result = q.FirstNumber + q.SecondNumber
}

func FindQuiz(id int) Quiz {
	for _, q := range quizes {
		if q.ID == id {
			return q
		}
	}
	return Quiz{}
}

func UpdateQuiz(q Quiz) error {
	for i, qi := range quizes {
		if qi.ID == q.ID {
			q.Score = q.Score
			q.FirstNumber = rand.Intn(10)
			q.SecondNumber = rand.Intn(10)
			q.CalcResult()
			q.Correct = false
			quizes[i] = q

			return nil
		}
	}
	return fmt.Errorf("Could not find Quiz with id of %d to delete", q.ID)
}
