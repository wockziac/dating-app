package usecase

import (
	"datingapp/application/core"
	"datingapp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitiateLogin(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		request := core.InitiateLoginRequest{
			PhoneNumber: "081234567891",
		}

		usrSessionRepo := new(mocks.ISessionRepository)
		usrSessionRepo.On("InitiateLogin", mock.Anything).Return(nil)

		usrCredRepo := new(mocks.ICredentialRepository)
		usrCredRepo.On("GetCredentials", mock.Anything).Return(&core.Credentials{
			UserID:       1,
			PhoneNumber:  request.PhoneNumber,
			EmailAddress: "banshee@gmng.com",
		}, nil)

		usrSessionUseCase := NewSessionUseCase(usrSessionRepo, usrCredRepo)
		err := usrSessionUseCase.InitiateLogin(request)

		assert.Nil(t, err)
	})
}
