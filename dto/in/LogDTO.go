package in

import (
	"errors"
	"fmt"
)

type UserRequest struct {
	Id        int64  `json:"-"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
}

func (input *UserRequest) ValidationRegistration() error {

	// ... validation logic here
	if input.FirstName == "" {
		return errors.New("first name is required")
	}

	if input.LastName == "" {
		return errors.New("last name is required")
	}

	if input.Gender != "L" || input.Gender != "P" {
		return errors.New("gender is required, Only L/P is allowed")

	}

	if input.Email != fmt.Sprintf("%s@%s", input.Username, "gmail.com") {
		return errors.New("email is required, use these format username@domain")
	}

	return input.mandatoryValidation()
}

func (input *UserRequest) mandatoryValidation() (err error) {
	if input.Username == "" {
		err = errors.New("Username is required")
		return
	}
	if input.Password == "" {
		err = errors.New("Password is required")
		return
	}

	return
}
