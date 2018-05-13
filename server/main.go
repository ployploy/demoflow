package main

import (
	"encoding/json"
	"io/ioutil"
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
	r.Handle("/user", CreateUserHandler).Methods("POST")

	//วิธีกำหนดคนเข้าถึงกรณีเรียกใช้งานข้าม IP
	corsObj := handlers.AllowedOrigins([]string{"*"})
	method := handlers.AllowedMethods([]string{"POST"})
	header := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})

	http.ListenAndServe(":3001", handlers.CORS(corsObj, method, header)(r))
}

// HealthHandler return api health
var HealthHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//เรียกใช้ service อื่นที่ต้องใช้งาน
	w.Write([]byte("API is up and running"))
})

var GetUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user := User{Username: "ploy", Email: "chonnikan@gmail.com"}
	response, _ := json.Marshal(user)

	//ให้ส่งข้อมูลกลับในรูปแบบ json
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
})

var CreateUserHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	response := "created user: " + user.Username + " email: " + user.Email

	w.Write([]byte(response))
})
