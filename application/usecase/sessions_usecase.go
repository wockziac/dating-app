package usecase

import (
	"datingapp/application/core"
	otputil "datingapp/application/utils/otp"
	"fmt"
)

type SessionUseCase struct {
	SessionRepo    core.ISessionRepository
	CredentialRepo core.ICredentialRepository
}

func NewSessionUseCase(usrsSessionRepo core.ISessionRepository, usrCredRepo core.ICredentialRepository) *SessionUseCase {
	return &SessionUseCase{
		SessionRepo:    usrsSessionRepo,
		CredentialRepo: usrCredRepo,
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

	return "", nil
}
