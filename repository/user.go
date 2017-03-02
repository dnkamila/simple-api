package repository

import (
	"gopkg.in/pg.v5"
	"gopkg.in/pg.v5/types"
	"log"
	"simple-api/database"
	. "simple-api/models"
)

func CreateUser(user *User) (*User, error) {
	err := database.GetDB().Insert(&User{
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

func GetUserById(user *User) (*User, error) {
	err := database.GetDB().Select(user)
	if err != nil {
		log.Println("Cannot retrieve user %v from database. Error: %s", user, err.Error())
		return nil, err
	}

	return user, nil
}
