package core

import (
	phonenumberutil "datingapp/application/utils/phoneNumber"
	"fmt"
	"time"
)

const (
	OTP_LENGTH = 4
)

type Session struct {
	Credential      Credentials
	SessionID       int
	OTP             string
	OTPExpiryDate   time.Time
	Token           string
	TokenExpiryDate time.Time
}

type ISessionUseCase interface {
	InitiateLogin(request InitiateLoginRequest) error
	Login(request LoginRequest) (string, error)
}

type ISessionRepository interface {
	InitiateLogin(session Session) error
	Login(session Session) (Session, error)
	GetSessionsByPhone(phoneNumber string) ([]Session, error)
}

type NewSessionRequest struct {
	Credential    Credentials
	OTP           string
	OTPExpiryDate time.Time
}

func NewSession(request NewSessionRequest) Session {
	return Session{
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
