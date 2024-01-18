package models

import "datingapp/application/core"

type RegisterCredentialsRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func (r *RegisterCredentialsRequest) ParseToCoreModel() core.CredentialRequest {
	return core.CredentialRequest{
		PhoneNumber:  r.PhoneNumber,
		EmailAddress: r.Email,
	}
}
