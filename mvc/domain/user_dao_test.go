package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//test case for non existing user
func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t, user, "we were not expecting a user with user id zero")
	//meaning of above statement is
	//if user == nil {
	// 	t.Error("we were not expecting a user with userId 0")
	// }

	assert.NotNil(t, err, "we were expecting error when user id is zero")
	// if err != nil {
	// 	t.Error("we were  expectiong error when user id is zero")
	// }

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)

	// if err.StatusCode == http.StatusNotFound {
	// 	t.Error("we are expecting 404 when user is not found")
	// }

}

//Test case for existing user
func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)

}
