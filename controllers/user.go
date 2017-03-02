package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "simple-api/models"
	"simple-api/repository"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)

	if err != nil {
		log.Println("Cannot decode JSON %v. Error: %s", r.Body, err.Error())
		return
	}

	insertedUser, err := repository.CreateUser(&user)
	if err != nil {
		log.Println("Cannot insert user")
		return
	}

	userJSON, err := json.Marshal(insertedUser)
	if err != nil {
		log.Println("Cannot encode JSON %v. Error: %s", insertedUser, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["id"])
	if vars["id"] == "" || err != nil {
		log.Println("Something wrong with id")
		return
	}
	user := User{Id: userId}

	retrievedUser, err := repository.GetUserById(&user)
	if err != nil {
		log.Println("Cannot retrive user")
		return
	}

	userJSON, err := json.Marshal(retrievedUser)
	if err != nil {
		log.Println("Cannot encode JSON %v. Error: %s", retrievedUser, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {

}
