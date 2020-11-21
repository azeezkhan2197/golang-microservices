package github_provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T){
	header := getAuthorizationHeader("123")
	assert.EqualValues(t,header,"token 123")
}
