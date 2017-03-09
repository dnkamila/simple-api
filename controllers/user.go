package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "simple-api/models"
	"simple-api/repository"
	"strconv"
	"fmt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)

	if err != nil {
		log.Printf("Cannot decode JSON %v. Error: %s", r.Body, err.Error())
		return
	}

	fmt.Printf("user %v\n", user)
	insertedUser, err := repository.GetUserRepository().CreateUser(&user)
	if err != nil {
		log.Println("Cannot insert user")
		return
	}

	userJSON, err := json.Marshal(insertedUser)
	if err != nil {
		log.Printf("Cannot encode JSON %v. Error: %s", insertedUser, err.Error())
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

	retrievedUser, err := repository.GetUserRepository().GetUserById(&user)
	if err != nil {
		log.Println("Cannot retrive user")
		return
	}

	userJSON, err := json.Marshal(retrievedUser)
	if err != nil {
		log.Printf("Cannot encode JSON %v. Error: %s", retrievedUser, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["username"] == "" {
		log.Println("Something wrong with username")
		return
	}
	user := User{Username: vars["username"]}

	retrievedUser, err := repository.GetUserRepository().GetUserByUsername(&user)
	if err != nil {
		log.Println("Cannot retrive user")
		return
	}

	userJSON, err := json.Marshal(retrievedUser)
	if err != nil {
		log.Printf("Cannot encode JSON %v. Error: %s", retrievedUser, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func UpdateUserPasswordById(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)

	if err != nil {
		log.Printf("Cannot decode JSON %v. Error: %s", r.Body, err.Error())
		return
	}

	updatedUser, err := repository.GetUserRepository().UpdateUserPasswordById(&user)
	if err != nil {
		log.Println("Cannot update user")
		return
	}

	userJSON, err := json.Marshal(updatedUser)
	if err != nil {
		log.Printf("Cannot encode JSON %v. Error: %s", updatedUser, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func UpdateUserPasswordByUsername(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	err := decoder.Decode(&user)

	if err != nil {
		log.Printf("Cannot decode JSON %v. Error: %s", r.Body, err.Error())
		return
	}

	updatedUser, err := repository.GetUserRepository().UpdateUserPasswordByUsername(&user)
	if err != nil {
		log.Println("Cannot update user")
		return
	}

	userJSON, err := json.Marshal(updatedUser)
	if err != nil {
		log.Printf("Cannot encode JSON %v. Error: %s", updatedUser, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJSON)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["id"])
	if vars["id"] == "" || err != nil {
		log.Println("Something wrong with id")
		return
	}
	user := User{Id: userId}

	err = repository.GetUserRepository().DeleteUserById(&user)
	if err != nil {
		log.Println("Cannot delete user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}

func DeleteUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["username"] == "" {
		log.Println("Something wrong with username")
		return
	}
	user := User{Username: vars["username"]}

	err := repository.GetUserRepository().DeleteUserByUsername(&user)
	if err != nil {
		log.Println("Cannot delete user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
