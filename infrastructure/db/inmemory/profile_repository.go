package inmemory

import (
	"datingapp/application/core"
	"sync"
)

type ProfileRepository struct {
	data     map[int]core.Profile
	writeMtx sync.RWMutex
}

func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{
		data:     make(map[int]core.Profile),
		writeMtx: sync.RWMutex{},
	}
}

func (r *ProfileRepository) GetUserProfile(profileID int) core.Profile {
	r.writeMtx.RLock()
	profile := r.data[profileID]
	defer r.writeMtx.RUnlock()
	return profile
}

func (r *ProfileRepository) InsertNewProfile(profile core.Profile) (core.Profile, error) {
	r.writeMtx.Lock()
	profile.ProfileID = len(r.data) + 1
	r.data[profile.ProfileID] = profile
	r.writeMtx.Unlock()

	return profile, nil
}
