package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data interface{}
	decoder.Decode(&data)

	fmt.Printf("[CreateToken] request body: %v\n", r.Body)
	fmt.Printf("[CreateToken] decoder: %v\n", decoder)
	fmt.Printf("[CreateToken] data: %v\n", data)
}