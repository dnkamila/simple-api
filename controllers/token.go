package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-api/helpers"
	. "simple-api/models"
	"simple-api/repository"
	"time"
)

func UpdateToken(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User
	decoder.Decode(&user)

	userRepository := repository.GetUserRepository()
	searchedUser, err := userRepository.GetUserById(&user)
	if err != nil {
		return
	}

	claimsSet := map[string]interface{}{
		"id":       searchedUser.Id,
		"username": searchedUser.Username,
	}
	token, err := helpers.CreateJWT(claimsSet, time.Now().Add(time.Minute*24*60))
	searchedUser.Token = token

	updatedUser, err := userRepository.UpdateUserTokenById(searchedUser)
	if err != nil {
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