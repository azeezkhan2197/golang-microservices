package github_provider

import (
	"errors"
	"github.com/azeezkhan2197/golang-microservices/src/api/client/restclient"
	"github.com/azeezkhan2197/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M){
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T){
	header := getAuthorizationHeader("123")
	assert.EqualValues(t,header,"token 123")
}

func TestCreateRepo(t *testing.T){
	//restclient.StartMockups()
	resp,err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,resp)
	assert.NotNil(t,err)
}

func TestCreateRepoErrorRestclient(t *testing.T){
	restclient.FlushMockups()
	restclient.AddMockUp(restclient.Mock{
		Url : "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err : errors.New("invalid rest client"),
	})
	response ,err :=CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,response)
	assert.NotNil(t,err)

	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)

	assert.EqualValues(t,"invalid rest client",err.Message)

	restclient.StopMockups()

}

func TestCreateRepoErrorInvalidResponseBody(t *testing.T){
	restclient.FlushMockups()
	invalidCloser , _ := os.Open("-asf3")
	restclient.AddMockUp(restclient.Mock{
		Url : "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode:http.StatusCreated,
			Body      :invalidCloser,
		},
	})
	response ,err :=CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,response)
	assert.NotNil(t,err)

	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)
	//assert.EqualValues(t,"invalid response body",err.Message)

	restclient.StopMockups()

}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error in json response body", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}



func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal", err.Message)
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()

	restclient.AddMockUp(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123,"name": "golang-tutorial","full_name": "azeezkhan2197/golang-tutorial"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)
	assert.EqualValues(t, "golang-tutorial", response.Name)
	assert.EqualValues(t, "azeezkhan2197/golang-tutorial", response.FullName)
}