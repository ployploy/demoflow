package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//1 routing per 1 handler
	r.Handle("/health", HealthHandler).Methods("GET")
	r.Handle("/users", GetUserHandler).Methods("GET")

	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

// HealthHandler return api health
var HealthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//เรียกใช้ service อื่นที่ต้องใช้งาน
	w.Write([]byte("API is up and running"))
})

var GetUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//เรียกใช้ service อื่นที่ต้องใช้งาน
	w.Write([]byte("chonnikan"))
})
