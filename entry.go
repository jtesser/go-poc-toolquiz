package main

import (
	"github.com/gorilla/mux"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/question", questionHandler).Methods("GET")
	headersOk := gorillaHandlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"})
	originsOk := gorillaHandlers.AllowedOrigins([]string{"*"})
	methodsOk := gorillaHandlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})

	http.ListenAndServe(":8080", gorillaHandlers.CORS(originsOk, headersOk, methodsOk)(router))
}

func questionHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte("helloworld"))
}
