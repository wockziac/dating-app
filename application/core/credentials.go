package core

import (
	phonenumberutil "datingapp/application/utils/phoneNumber"
	"fmt"
	"regexp"
)

type Credentials struct {
	UserID       int
	PhoneNumber  string
	EmailAddress string
}

type CredentialRequest struct {
	PhoneNumber  string
	EmailAddress string
}

type ICredentialRepository interface {
	GetCredentials(phoneNumber string) (*Credentials, error)
	InsertCredentials(Credentials) (Credentials, error)
}

type ICredentialUseCase interface {
	RegisterCredential(request CredentialRequest) error
}

const (
	EMAIL_REGEX = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

func CreateNewCredential(request CredentialRequest) (Credentials, error) {
	uc := Credentials{
		PhoneNumber:  request.PhoneNumber,
		EmailAddress: request.EmailAddress,
	}

	err := uc.Validate()
	return uc, err
}

func (uc *Credentials) Validate() error {
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
