package tests

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"simple-api/application"
	"simple-api/controllers"
	. "simple-api/models"
	"simple-api/repository"
	"strconv"
	"strings"
	"testing"
)

const (
	basedUserURL         = "/user"
	basedUserIdUrl       = "/user/id"
	basedUserUsernameUrl = "/user/username"
)

var (
	id         = 1
	username   = "kamila@icehousecorp.com"
	password   = "12345"
	newPasword = "54321"
	user       = User{
		Id:       id,
		Username: username,
		Password: password,
	}
	newUser = User{
		Id:       id,
		Username: username,
		Password: newPasword,
	}
)

func TestCreateUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Id:       0,
		Username: "kamila@icehousecorp.com",
		Password: "12345",
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().CreateUser(&inputUser).Return(&user, nil)
	repository.SetUserRepository(mockUserRepository)

	jsonRequest := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, user.Username, user.Password)

	w := httptest.NewRecorder()

	r := httptest.NewRequest("POST", basedUserURL, strings.NewReader(jsonRequest))
	r.Header.Set("Content-Type", "application/json")
	controllers.CreateUser(w, r)
}

func TestGetUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Id: user.Id,
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().GetUserById(&inputUser).Return(&user, nil)
	repository.SetUserRepository(mockUserRepository)

	app := application.NewApp()
	app.InitRouter()

	server := httptest.NewServer(app.Router)
	defer server.Close()

	url := fmt.Sprintf("%s%s/%s", server.URL, basedUserIdUrl, strconv.Itoa(user.Id))
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("unsuccess with status code: %d", res.StatusCode)
	}
}

func TestGetUserByUsername(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Username: user.Username,
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().GetUserByUsername(&inputUser).Return(&user, nil)
	repository.SetUserRepository(mockUserRepository)

	app := application.NewApp()
	app.InitRouter()

	server := httptest.NewServer(app.Router)
	defer server.Close()

	url := fmt.Sprintf("%s%s/%s", server.URL, basedUserUsernameUrl, user.Username)
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("unsuccess with status code: %d", res.StatusCode)
	}
}

func TestUpdateUserPasswordById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Id:       user.Id,
		Password: newPasword,
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().UpdateUserPasswordById(&inputUser).Return(&newUser, nil)
	repository.SetUserRepository(mockUserRepository)

	jsonRequest := fmt.Sprintf(`{"id": %d, "password": "%s"}`, user.Id, newUser.Password)

	w := httptest.NewRecorder()

	r := httptest.NewRequest("PUT", basedUserIdUrl, strings.NewReader(jsonRequest))
	r.Header.Set("Content-Type", "application/json")
	controllers.UpdateUserPasswordById(w, r)
}

func TestUpdateUserPasswordByUsername(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Username: user.Username,
		Password: newPasword,
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().UpdateUserPasswordByUsername(&inputUser).Return(&newUser, nil)
	repository.SetUserRepository(mockUserRepository)

	jsonRequest := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, user.Username, newUser.Password)

	w := httptest.NewRecorder()

	r := httptest.NewRequest("PUT", basedUserIdUrl, strings.NewReader(jsonRequest))
	r.Header.Set("Content-Type", "application/json")
	controllers.UpdateUserPasswordByUsername(w, r)
}

func TestDeleteUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Id: user.Id,
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().DeleteUserById(&inputUser).Return(nil)
	repository.SetUserRepository(mockUserRepository)

	app := application.NewApp()
	app.InitRouter()

	server := httptest.NewServer(app.Router)
	defer server.Close()

	url := fmt.Sprintf("%s%s/%s", server.URL, basedUserIdUrl, strconv.Itoa(user.Id))
	req, err := http.NewRequest("DELETE", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("unsuccess with status code: %d", res.StatusCode)
	}
}

func TestDeleteUserByUsername(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	inputUser := User{
		Username: user.Username,
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().DeleteUserByUsername(&inputUser).Return(nil)
	repository.SetUserRepository(mockUserRepository)

	app := application.NewApp()
	app.InitRouter()

	server := httptest.NewServer(app.Router)
	defer server.Close()

	url := fmt.Sprintf("%s%s/%s", server.URL, basedUserUsernameUrl, user.Username)
	req, err := http.NewRequest("DELETE", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("unsuccess with status code: %d", res.StatusCode)
	}
}
