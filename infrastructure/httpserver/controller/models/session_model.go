package models

import "datingapp/application/core"

type InitiateLoginRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

type CompleteLoginRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	OTP         string `json:"otp"`
}

func (r *CompleteLoginRequest) ParseToCoreModel() core.LoginRequest {
	return core.LoginRequest{
		PhoneNumber: r.PhoneNumber,
		OTP:         r.OTP,
	}
}

type CompleteLoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
