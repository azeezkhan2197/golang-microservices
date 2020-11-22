package  restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)


var (
	enableMocks = false
	mocks = make(map[string]*Mock)
)



type Mock struct{
	Url string
	HttpMethod string
	Response *http.Response
	Err error
}

func getMockId(httpMethod string,url string ) string{
	return fmt.Sprintf("%s_%s",httpMethod,url)
}


func StartMockups(){
	enableMocks = true
}

func StopMockups(){
	enableMocks = false
}

func FlushMockups(){
	mocks = make(map[string]*Mock)
}
func AddMockUp(mock Mock){
	mocks[getMockId(mock.HttpMethod,mock.Url)] = &mock
}



func Post(url string,body interface{},headers http.Header) (*http.Response,error){
	if enableMocks{
		//return local mocks without calling an external resource

		mock := mocks[getMockId(http.MethodPost,url)]
		if mock == nil{
			return nil,errors.New("no mock up found for given request")
		}
		return mock.Response,mock.Err
		}

	jsonBytes,err := json.Marshal(body)
	if err!=nil{
		return nil,err
	}
	request,err := http.NewRequest(http.MethodPost,url,bytes.NewReader(jsonBytes))
	request.Header = headers
	client := http.Client{}
	return client.Do(request)

}