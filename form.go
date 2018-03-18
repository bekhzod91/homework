package main

import "strings"

type UserFormErrors struct {
	Email string
	FirstName string
	SecondName string
}

func UserForm(data *RequestUserData) (bool, *UserFormErrors) {
	errors := new(UserFormErrors)

	if !strings.Contains(data.Email, "@") {
		errors.Email = "Enter a valid email address."
	}

	if data.FirstName == "" {
		errors.FirstName = "This field may not be blank."
	}

	if data.SecondName == "" {
		errors.SecondName = "This field may not be blank."
	}

	isValid := (UserFormErrors{}) == *errors

	return isValid, errors
}
