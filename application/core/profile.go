package core

import (
	"time"
)

type Profile struct {
	UserID      int
	ProfileID   int
	Name        string
	DOB         time.Time
	Gender      string
	Interests   Interest
	Description string
}

type Interest struct {
	Hobby     []string
	Interests []string
}

type IProfileRepository interface {
	GetUserProfile(profileID int) (Profile, error)
	InsertNewProfile(profile Profile) (Profile, error)
}

type IProfileUseCase interface {
	RegisterProfile(profile Profile) error
}
