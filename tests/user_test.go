package tests

import (
	"testing"
	"github.com/golang/mock/gomock"
	. "simple-api/models"
	"fmt"
	"net/http/httptest"
	"strings"
	"simple-api/application"
	"simple-api/repository"
	"simple-api/controllers"
	"strconv"
	"net/http"
)

const (
	basedUserURL = "/user"
	basedUserIdUrl = "/user/id"
	basedUserUsernameUrl = "/user/username"
)

var (
	id = 1
	username = "kamila@icehousecorp.com"
	password = "12345"
	newPasword = "54321"
	user = User{
		Id: id,
		Username: username,
		Password: password,
	}
	newUser = User{
		Id: id,
		Username: username,
		Password: newPasword,
	}

)


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

	url := fmt.Sprintf("%s/%s/%s", server.URL, basedUserIdUrl, strconv.Itoa(user.Id))
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

	url := fmt.Sprintf("%s/%s/%s", server.URL, basedUserUsernameUrl, user.Username)
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
		Id: user.Id,
		Password: newPasword,
	}
	mockUserRepository := repository.NewMockUserRepositoryInterface(mockCtrl)
	mockUserRepository.EXPECT().UpdateUserPasswordById(&inputUser).Return(&newUser, nil)
	repository.SetUserRepository(mockUserRepository)

	fmt.Println("before")
	jsonRequest := fmt.Sprintf(`{"id": %d, "password": "%s"}`, user.Id, newUser.Password)
	fmt.Println("after")

	w := httptest.NewRecorder()

	r := httptest.NewRequest("PUT", basedUserIdUrl, strings.NewReader(jsonRequest))
	r.Header.Set("Content-Type", "application/json")
	controllers.UpdateUserPasswordById(w, r)
}
