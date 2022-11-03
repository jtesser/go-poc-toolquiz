package main

import (
	"encoding/json"
	"fmt"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-poc-toolquiz/quiz"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/question/random", randomQuestionHandler).Methods("GET")
	router.HandleFunc("/api/question/hint/{id}", hintHandler).Methods("GET")
	router.HandleFunc("/api/question/answer/{id}/{answer}", answerHandler).Methods("GET")

	headersOk := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"})
	originsOk := gorillaHandlers.AllowedOrigins([]string{"*"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{http.MethodGet})

	sTest := "mystring"
	rTest := ""
	for i := 0; i < len(sTest); i++ {
		rTest = string(sTest[i]) + rTest
	}
	fmt.Println(rTest)

	http.ListenAndServe(":8080", gorillaHandlers.CORS(originsOk, headersOk, methodsOk)(router))
}

func randomQuestionHandler(writer http.ResponseWriter, request *http.Request) {
	t, err := quiz.Random()
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	ret, err := json.Marshal(t)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	_, err = writer.Write(ret)
	if err != nil {
		fmt.Print("error is ", err)
	}
}

func hintHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("id is not valid")
		writer.WriteHeader(500)
		return
	}

	hint, err := quiz.Hint(id)
	writer.Header().Set("Content-Type", "text/plain")
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	_, err = writer.Write([]byte(hint))
	if err != nil {
		fmt.Print("error is ", err)
	}
}

func answerHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		writer.WriteHeader(500)
		return
	}

	answer := params["answer"]
	if len(answer) < 1 {
		fmt.Println("answer is not valid")
		writer.WriteHeader(500)
		return
	}

	correct, err := quiz.Answer(id, answer)
	writer.Header().Set("Content-Type", "text/plain")
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	if correct {
		_, err = writer.Write([]byte("true"))
	} else {
		_, err = writer.Write([]byte("false"))
	}

	if err != nil {
		fmt.Print("error is ", err)
	}
}
