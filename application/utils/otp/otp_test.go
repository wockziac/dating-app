package otputil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {
	currentTime := time.Now()
	otp, expiryTime := GenerateOTP()
	assert.NotEqual(t, "", otp)

	timeDifference := expiryTime.Sub(currentTime)
	assert.Equal(t, 45, int(timeDifference.Abs().Seconds()))
}
