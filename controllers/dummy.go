package controllers

import (
	"net/http"
	"simple-api/helpers"
	"fmt"
)

func Dummy(w http.ResponseWriter, r *http.Request) {
	headerAuth := r.Header.Get("Authorization")
	fmt.Printf("headerAuth: %v\n", headerAuth)
	token, _ := helpers.ParseHeaderJWT(headerAuth)
	fmt.Printf("token: %v\n", token)
	claims, _ := helpers.ValidateJWT(token)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"isValid"}`))

	fmt.Printf("claims: %v\n", claims)
}