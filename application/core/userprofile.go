package core

import (
	"time"
)

type UserProfile struct {
	UserID        int
	UserProfileID int
	Name          string
	DOB           time.Time
	Gender        string
	Interests     UserInterest
	Description   string
}

type UserInterest struct {
	Hobby     []string
	Interests []string
}

type IUserProfileRepository interface {
	GetUserProfile(profileID int) (UserProfile, error)
	InsertNewProfile(userProfile UserProfile) (UserProfile, error)
}

type IUserProfileUseCase interface {
	RegisterUserProfile(userProfile UserProfile) error
}
