package models

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeLogin(t *testing.T) {
	f := `
{
"username": "asimbera",
"password": "some_password"
}
	`
	a := LoginInput{Username: "asimbera", Password: "some_password"}
	var u LoginInput

	err := json.Unmarshal([]byte(f), &u)
	assert.Equal(t, err, nil, "Decoding shouldn't return error.")

	validate := validator.New()
	err = validate.Struct(u)
	assert.Equal(t, err, nil, "Validation shouldn't return error.")

	assert.Equal(t, a, u, "Decoded and provided data should be same.")
}

func TestDecodeSignup(t *testing.T) {
	fOne := `
{
	"name": "Asim Bera",
	"email": "mail@example.com",
	"username": "asimbera",
	"password": "some_password"
}
`

	fTwo := `
{
	"name": "Asim Bera",
	"email": "mail@example.com",
	"username": "asimbera",
	"password": "some_password"
}
`

	a := SignupInput{
		Name:     "Asim Bera",
		Email:    "mail@example.com",
		Username: "asimbera",
		Password: "some_password",
	}

	var uOne, uTwo SignupInput

	err := json.Unmarshal([]byte(fOne), &uOne)
	assert.Equal(t, err, nil, "Decoding shouldn't return error.")

	err = json.Unmarshal([]byte(fTwo), &uTwo)
	assert.Equal(t, err, nil, "Decoding shouldn't return error.")

	validate := validator.New()
	err = validate.Struct(uOne)
	assert.Equal(t, err, nil, "Validation shouldn't return error.")

	assert.Equal(t, a, uOne, "Decoded and provided data should be same.")
	assert.NotEqual(t, a, uTwo, "Decoded and provided data shouldn't be same.")

}
