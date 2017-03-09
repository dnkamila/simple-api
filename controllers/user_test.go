package controllers

import (
	"testing"
	"simple-api/repository"
	"github.com/golang/mock/gomock"
	. "simple-api/models"
	"fmt"
	"net/http/httptest"
	"strings"
)

const (
	createUserURL = "/user"
)

var user = User{
	Id: 1,
	Username: "kamila@icehousecorp.com",
	Password: "12345",
}

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Id: 0,
		Username: "kamila@icehousecorp.com",
		Password: "12345",
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().CreateUser(&inputUser).Return(&user, nil)
	repository.SetUserRepository(mockUserRepository)

	jsonRequest := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, user.Username, user.Password)

	w := httptest.NewRecorder()

	r := httptest.NewRequest("POST", createUserURL, strings.NewReader(jsonRequest))
	r.Header.Set("Content-Type", "application/json")

	CreateUser(w, r)
}