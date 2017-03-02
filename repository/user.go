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
		log.Printf("Cannot insert user %v to database. Error: %s", user, err.Error())
		return nil, err
	}

	_, id, err := getLastInsertUserId()
	if err != nil {
		log.Printf("Cannot get last insert id. Error: %s", err.Error())
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
		log.Printf("Cannot retrieve user %v from database. Error: %s", user, err.Error())
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(user *User) (*User, error) {
	err := database.GetDB().Model(&user).Where("username = ?", user.Username).Select()
	if err != nil {
		log.Printf("Cannot retrieve user %v from database. Error: %s", user, err.Error())
		return nil, err
	}

	return user, nil
}

func UpdateUserPasswordById(user *User) (*User, error) {
	_, err := database.GetDB().Model(&user).Column("password").Returning("*").Update()
	if err != nil {
		log.Printf("Cannot update user %v in database. Error: %s", user, err.Error())
		return nil, err
	}

	return GetUserById(user)
}