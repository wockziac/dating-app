package core

import (
	"datingapp/application/utils/date"
	phonenumberutil "datingapp/application/utils/phoneNumber"
	"fmt"
)

const (
	OTP_LENGTH = 4
)

type UserSession struct {
	UserID          UserCredentials
	SessionID       int
	OTP             string
	OTPExpiryDate   date.Date
	Token           string
	TokenExpiryDate date.Date
}

type IUserSessionUseCase interface {
	InitiateLogin(request InitiateLoginRequest)
	Login(request LoginRequest) UserSession
}

type IUserSessionRepository interface {
	InitiateLogin(session UserSession) error
	Login(session UserSession) (UserSession, error)
	GetUserSessionsByPhone(phoneNumber string) ([]UserSession, error)
}

type InitiateLoginRequest struct {
	PhoneNumber string
	// Other attributes can be added such as device, IP, etc
}

func (r *InitiateLoginRequest) Validate() error {
	number, err := phonenumberutil.ParsePhoneNumber(r.PhoneNumber)
	if err != nil {
		return err
	}
	r.PhoneNumber = number

	return nil
}

type LoginRequest struct {
	PhoneNumber string
	OTP         string
}

func (r *LoginRequest) Validate() error {
	number, err := phonenumberutil.ParsePhoneNumber(r.PhoneNumber)
	if err != nil {
		return err
	}
	r.PhoneNumber = number

	if r.OTP == "" || len(r.OTP) != OTP_LENGTH {
		return fmt.Errorf("incorrect OTP")
	}

	return nil
}
