package repository

import (
	"log"
	"simple-api/database"
	. "simple-api/models"
	"gopkg.in/pg.v5/types"
	"gopkg.in/pg.v5"
)

func CreateUser(user *User) (*User, error) {
	err := database.GetDB().Insert(&User {
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		log.Println("Cannot insert user %v to database. Error: %s", user, err.Error())
		return nil, err
	}

	_, id, err := getLastInsertUserId()
	if err != nil {
		log.Println("Cannot get last insert id. Error: %s", err.Error())
		return nil, err
	}

	user.Id = id

	return user, nil
}

func getLastInsertUserId() (*types.Result, int, error) {
	var id int
	result, err := database.GetDB().QueryOne(pg.Scan(&id), `SELECT currval('users_id_seq')`)

	return result, id, err
}