package repository

import (
	"gopkg.in/pg.v5"
	"gopkg.in/pg.v5/types"
	"log"
	"simple-api/database"
	. "simple-api/models"
)

type UserRepositoryInterface interface {
	CreateUser(user *User) (*User, error)
	GetUserById(user *User) (*User, error)
	GetUserByUsername(user *User) (*User, error)
	UpdateUserPasswordById(user *User) (*User, error)
	UpdateUserPasswordByUsername(user *User) (*User, error)
	UpdateUserTokenById(user *User) (*User, error)
	DeleteUserById(user *User) error
	DeleteUserByUsername(user *User) error
}

var userRepository UserRepositoryInterface

type UserRepository struct{}

func GetUserRepository() UserRepositoryInterface {
	if userRepository == nil {
		return &UserRepository{}
	}

	return userRepository
}

func SetUserRepository(instance UserRepositoryInterface) {
	userRepository = instance
}

func (*UserRepository) CreateUser(user *User) (*User, error) {
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

func (*UserRepository) GetUserById(user *User) (*User, error) {
	err := database.GetDB().Select(user)
	if err != nil {
		log.Printf("Cannot retrieve user %v from database. Error: %s", user, err.Error())
		return nil, err
	}

	return user, nil
}

func (*UserRepository) GetUserByUsername(user *User) (*User, error) {
	err := database.GetDB().Model(&user).Where("username = ?", user.Username).Select()
	if err != nil {
		log.Printf("Cannot retrieve user %v from database. Error: %s", user, err.Error())
		return nil, err
	}

	return user, nil
}

func (this *UserRepository) UpdateUserPasswordById(user *User) (*User, error) {
	_, err := database.GetDB().Model(&user).Column("password").Returning("*").Update()
	if err != nil {
		log.Printf("Cannot update user %v in database. Error: %s", user, err.Error())
		return nil, err
	}

	return this.GetUserById(user)
}

func (this *UserRepository) UpdateUserPasswordByUsername(user *User) (*User, error) {
	_, err := database.GetDB().Model(&user).
		Set("password = ?", user.Password).
		Where("username = ?", user.Username).
		Returning("*").
		Update()
	if err != nil {
		log.Printf("Cannot update user %v in database. Error: %s", user, err.Error())
		return nil, err
	}

	return this.GetUserById(user)
}

func (this *UserRepository) UpdateUserTokenById(user *User) (*User, error) {
	_, err := database.GetDB().Model(&user).Column("token").Returning("*").Update()
	if err != nil {
		log.Printf("Cannot update user %v in database. Error: %s", user, err.Error())
		return nil, err
	}

	return this.GetUserById(user)
}

func (*UserRepository) DeleteUserById(user *User) error {
	return database.GetDB().Delete(user)
}

func (*UserRepository) DeleteUserByUsername(user *User) error {
	_, err := database.GetDB().Model(&user).Where("username = ?", user.Username).Delete()
	return err
}
