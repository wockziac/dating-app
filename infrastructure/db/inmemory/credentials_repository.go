package inmemory

import (
	"datingapp/application/core"
	"sync"
)

type CredentialRepository struct {
	data     map[string]core.Credentials
	writeMtx sync.RWMutex
}

func NewCredentialRepository() *CredentialRepository {
	return &CredentialRepository{
		data:     make(map[string]core.Credentials),
		writeMtx: sync.RWMutex{},
	}
}

func (r *CredentialRepository) GetCredentials(phoneNumber string) (*core.Credentials, error) {
	r.writeMtx.RLock()
	defer r.writeMtx.RUnlock()
	data, exist := r.data[phoneNumber]

	if !exist {
		return nil, nil
	}

	return &core.Credentials{
		UserID:       data.UserID,
		PhoneNumber:  data.PhoneNumber,
		EmailAddress: data.EmailAddress,
	}, nil
}

func (r *CredentialRepository) InsertCredentials(credential core.Credentials) (core.Credentials, error) {
	r.writeMtx.Lock()
	credential.UserID = len(r.data) + 1
	r.data[credential.PhoneNumber] = credential
	r.writeMtx.Unlock()
	return credential, nil
}
