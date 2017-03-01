package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	. "simple-api/models"
	"simple-api/repository"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)

	if err != nil {
		log.Println("Cannot decode JSON %v. Error: %s", r.Body, err.Error())
	}

	insertedUser, err := repository.CreateUser(&user)
	if err != nil {
		log.Println("Cannot insert user")
	}

	userJSON, err := json.Marshal(insertedUser)
	if err != nil {
		log.Println("Cannot encode JSON %v. Error: %s", insertedUser, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}
