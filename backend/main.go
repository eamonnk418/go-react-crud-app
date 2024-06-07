package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type User struct {
    ID          *string `json:"id"`
    FirstName   *string `json:"firstName"`
    MiddleName  *string `json:"middleName"`
    LastName    *string `json:"lastName"`
    Email       *string `json:"email"`
    Gender      *string `json:"gender"`
    CivilStatus *string `json:"civilStatus"`
    Birthday    *string `json:"birthday"`
    Contact     *string `json:"contact"`
    Address     *string `json:"address"`
    Age         *string `json:"age"`
}

var users = []User{
	{
		ID: uuid,
	}
}

func listAllUsers(w http.ResponseWriter, r *http.Request) {
	var userList []User
	for _, user := range users {
		userList = append(userList, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userList)
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /users", listAllUsers)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
