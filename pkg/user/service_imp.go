package user

import (
	"ss/pkg/crypt"
)

// ServiceImp provides simple and secured implementation for Service.
type ServiceImp struct {
	Repository Repository
	PwdCrypt   crypt.Service
}

// CreateNewUser persists user's ID and hashed password. Need to check for existing user with HasUser.
func (u *ServiceImp) CreateNewUser(id string, pwd string) error {
	return u.Repository.Create(&Model{
		ID:       id,
		Password: u.PwdCrypt.Hash(pwd),
	})
}

// HasUser returns true if there is a user with the given id.
func (u *ServiceImp) HasUser(id string) (bool, error) {
	user, err := u.Repository.Find(id)
	return user != nil, err
}

// Verify returns true if user ID and password is found in the repository.
func (u *ServiceImp) Verify(id string, pwd string) (bool, error) {
	user, err := u.Repository.Find(id)
	if err != nil {
		return false, err
	}

	return user != nil && u.PwdCrypt.Match(user.Password, pwd), nil
}
