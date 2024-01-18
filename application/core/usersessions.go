package core

import (
	phonenumberutil "datingapp/application/utils/phoneNumber"
	"fmt"
	"time"
)

const (
	OTP_LENGTH = 4
)

type UserSession struct {
	Credential      UserCredentials
	SessionID       int
	OTP             string
	OTPExpiryDate   time.Time
	Token           string
	TokenExpiryDate time.Time
}

type IUserSessionUseCase interface {
	InitiateLogin(request InitiateLoginRequest) error
	Login(request LoginRequest) (string, error)
}

type IUserSessionRepository interface {
	InitiateLogin(session UserSession) error
	Login(session UserSession) (UserSession, error)
	GetUserSessionsByPhone(phoneNumber string) ([]UserSession, error)
}

type NewUserSessionRequest struct {
	Credential    UserCredentials
	OTP           string
	OTPExpiryDate time.Time
}

func NewUserSession(request NewUserSessionRequest) UserSession {
	return UserSession{
		Credential:    request.Credential,
		OTP:           request.OTP,
		OTPExpiryDate: request.OTPExpiryDate,
	}
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
