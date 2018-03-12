package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		pattern := `%s - "%s %s %s %s"`

		log.Printf(pattern, r.RemoteAddr, r.Proto, r.Method, r.RequestURI, time.Since(start))

		next.ServeHTTP(w, r)
	})
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	//	Routes()
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/new", StartQuiz)
	mux.HandleFunc("/quizes", QuizIndex)
	mux.HandleFunc("/quizes/", QuizIndex)
	mux.HandleFunc("/quiz/", MyQuiz)

	log.Println("listener : Started : Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", LogRequest(mux)))
}
