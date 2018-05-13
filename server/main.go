package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	r := mux.NewRouter()

	//1 routing per 1 handler
	r.Handle("/health", HealthHandler).Methods("GET")
	r.Handle("/user", GetUserHandler).Methods("GET")

	//วิธีกำหนดคนเข้าถึงกรณีเรียกใช้งานข้าม IP
	corsObj := handlers.AllowedOrigins([]string{"*"})

	http.ListenAndServe(":3001", handlers.CORS(corsObj)(r))
}

// HealthHandler return api health
var HealthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//เรียกใช้ service อื่นที่ต้องใช้งาน
	w.Write([]byte("API is up and running"))
})

var GetUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user := User{Username: "ploy", Email: "chonnikan@gmail.com"}
	response, _ := json.Marshal(user)
	w.Write([]byte(response))
})
