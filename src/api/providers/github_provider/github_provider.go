package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/azeezkhan2197/golang-microservices/src/api/client/restclient"
	"github.com/azeezkhan2197/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const(
	headerAuthorization  		=  	"Authorization"
	headerAuthorizationFormat 	= 	"token %s"
	urlCreateRepo				=   "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string ) string{
	return fmt.Sprintf(headerAuthorizationFormat,accessToken)
}
func CreateRepo(accessToken string ,request github.CreateRepoRequest)(*github.CreateRepoResponse , *github.GitHubErrorResponse){
	//Header format is => Authorization  : token 92ebdb216babcc2660c6d76330706fcbdb2a64ea

	headers := http.Header{}
	headers.Set(headerAuthorization,getAuthorizationHeader(accessToken))
	response,err := restclient.Post(urlCreateRepo,request,headers)
	fmt.Println(response)
	//fmt.Println(err.Error())
	if err!=nil{
		log.Println(fmt.Sprintf("error while creating a new repo in github %s",err.Error()))
		return nil,&github.GitHubErrorResponse{
			StatusCode :  http.StatusInternalServerError,
			Message    :  err.Error(),
		}
	}
	//Read all the response
	byte,err:= ioutil.ReadAll(response.Body)
	if err!=nil{
		return nil,&github.GitHubErrorResponse{StatusCode:http.StatusInternalServerError,Message:"invalid response body"}
	}
	defer response.Body.Close()

	if response.StatusCode>299{
		var errResponse github.GitHubErrorResponse
		//convert the json response to struct
		if err:= json.Unmarshal(byte,&errResponse); err!=nil{
			return nil,&github.GitHubErrorResponse{StatusCode:http.StatusInternalServerError,Message:"error in json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil,&errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(byte,result);err!=nil{
		log.Println(fmt.Sprintf("error while trying to unmarshall  %s",err.Error()))
		return nil,&github.GitHubErrorResponse{StatusCode:http.StatusInternalServerError,Message:"error when trying to unmarshal"}

	}

	return &result,nil

}