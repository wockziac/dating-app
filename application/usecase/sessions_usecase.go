package usecase

import (
	"datingapp/application/core"
	otputil "datingapp/application/utils/otp"
	"fmt"
	"time"
)

type SessionUseCase struct {
	SessionRepo          core.ISessionRepository
	CredentialRepo       core.ICredentialRepository
	SessionTokenProvider core.ISessionTokenProvider
}

func NewSessionUseCase(usrsSessionRepo core.ISessionRepository, usrCredRepo core.ICredentialRepository, tokenProvider core.ISessionTokenProvider) *SessionUseCase {
	return &SessionUseCase{
		SessionRepo:          usrsSessionRepo,
		CredentialRepo:       usrCredRepo,
		SessionTokenProvider: tokenProvider,
	}
}

func (uc *SessionUseCase) InitiateLogin(request core.InitiateLoginRequest) error {

	usrCrd, err := uc.CredentialRepo.GetCredentials(request.PhoneNumber)
	if err != nil {
		return fmt.Errorf("error when trying to find user")
	}

	otp, expiryTime := otputil.GenerateOTP()
	usrSession := core.NewSession(core.NewSessionRequest{
		Credential:    *usrCrd,
		OTP:           otp,
		OTPExpiryDate: expiryTime,
	})
	err = uc.SessionRepo.InitiateLogin(usrSession)
	if err != nil {
		return fmt.Errorf("error initiate user login")
	}

	fmt.Println(fmt.Sprintf("OTP: %v. ExpiryTime: %v", otp, expiryTime))

	return nil
}

func (uc *SessionUseCase) Login(request core.LoginRequest) (string, error) {

	sessions, err := uc.SessionRepo.GetSessionsByPhone(request.PhoneNumber)
	if err != nil {
		return "", fmt.Errorf("error when get user sessions")
	}

	var session *core.Session
	for _, item := range sessions {
		otpExpiration := item.OTPExpiryDate.Sub(time.Now()).Seconds()
		if item.OTP == request.OTP && otpExpiration > 0 {
			session = item
			break
		}
	}

	if session == nil {
		return "", fmt.Errorf("failed login")
	}

	sessionToken, err := uc.SessionTokenProvider.GenerateToken(*session)
	session.UpdateSessionToken(sessionToken)
	if err != nil {
		return "", fmt.Errorf("error generate token")
	}

	_, err = uc.SessionRepo.Login(session)
	if err != nil {
		return "", fmt.Errorf("error store session token")
	}

	return sessionToken.Token, nil
}
