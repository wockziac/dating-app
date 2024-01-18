package inmemory

import (
	"datingapp/application/core"
	"sync"
)

type SessionRepository struct {
	data     map[int]core.Session
	writeMtx sync.RWMutex
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{
		data:     make(map[int]core.Session),
		writeMtx: sync.RWMutex{},
	}
}

func (s *SessionRepository) InitiateLogin(session *core.Session) error {
	s.writeMtx.Lock()
	session.SessionID = len(s.data) + 1
	s.data[session.SessionID] = *session
	s.writeMtx.Unlock()

	return nil
}

func (s *SessionRepository) Login(session *core.Session) (*core.Session, error) {
	s.writeMtx.Lock()
	s.data[session.SessionID] = *session
	s.writeMtx.Unlock()

	return session, nil
}

func (s *SessionRepository) GetSessionsByPhone(phoneNumber string) ([]*core.Session, error) {
	s.writeMtx.RLock()
	defer s.writeMtx.RUnlock()

	result := []*core.Session{}
	for _, item := range s.data {
		if item.Credential.PhoneNumber == phoneNumber {
			result = append(result, &core.Session{
				Credential:      item.Credential,
				SessionID:       item.SessionID,
				OTP:             item.OTP,
				OTPExpiryDate:   item.OTPExpiryDate,
				Token:           item.Token,
				TokenExpiryDate: item.TokenExpiryDate,
			})
		}
	}

	return result, nil
}
