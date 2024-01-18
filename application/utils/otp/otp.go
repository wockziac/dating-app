package otputil

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	OTP_LENGTH         = 6
	OTP_TIMEOUT_SECOND = 45
)

func GenerateOTP() (string, time.Time) {
	rand.Seed(time.Now().UnixNano())

	min := int64(1)
	max := int64(1)
	for i := 0; i < OTP_LENGTH; i++ {
		max *= 10
	}

	expiryTime := time.Now().Add(time.Second * OTP_TIMEOUT_SECOND)

	return fmt.Sprint(min + rand.Int63n(max-min)), expiryTime
}
