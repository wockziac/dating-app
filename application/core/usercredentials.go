package core

import (
	phonenumberutil "datingapp/application/utils/phoneNumber"
	"fmt"
	"regexp"
)

type UserCredentials struct {
	UserID       int
	PhoneNumber  string
	EmailAddress string
}

type UserCredentialRequest struct {
	PhoneNumber  string
	EmailAddress string
}

type IUserCredentialRepository interface {
	GetUserCredentials(phoneNumber string) (*UserCredentials, error)
	InsertUserCredentials(UserCredentials) (UserCredentials, error)
}

type IUserCredentialUseCase interface {
	// GetUserCredentials(phoneNumber, email string) (string, error)
	RegisterUserCredential(request UserCredentialRequest) error
}

const (
	EMAIL_REGEX = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func CreateNewUserCredential(request UserCredentialRequest) (UserCredentials, error) {
	uc := UserCredentials{
		PhoneNumber:  request.PhoneNumber,
		EmailAddress: request.EmailAddress,
	}

	err := uc.Validate()
	return uc, err
}

func (uc *UserCredentials) Validate() error {
	if uc.EmailAddress == "" {
		return fmt.Errorf("empty email address")
	}

	regex := regexp.MustCompile(EMAIL_REGEX)
	if !regex.MatchString(uc.EmailAddress) {
		return fmt.Errorf("invalid email address")
	}

	if uc.PhoneNumber == "" {
		return fmt.Errorf("empty phone number")
	}

	number, err := phonenumberutil.ParsePhoneNumber(uc.PhoneNumber)
	if err != nil {
		return err
	}

	uc.PhoneNumber = number

	return nil
}
