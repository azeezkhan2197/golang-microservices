package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepo(t *testing.T) {
	request := CreateRepo{
		Name:        "go-lang-introduction",
		Description: "go lang introdcution repository",
		HomePage:    "https://github.com",
		Private:     false,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, request)
	fmt.Println(string(bytes))

	var target CreateRepo
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
}
