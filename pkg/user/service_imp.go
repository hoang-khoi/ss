package user

import (
	"ss/pkg/pwdcrypt"
)

// ServiceImp provides simple and secured implementation for Service.
type ServiceImp struct {
	Repository Repository
	PwdCrypt   pwdcrypt.Service
}

// CreateNewUser persists user's ID and hashed password. Need to check for existing user with HasUser.
func (u *ServiceImp) CreateNewUser(m *Model) error {
	m.Password = u.PwdCrypt.Hash(m.Password)
	return u.Repository.Create(m)
}

// HasUser returns true if there is a user with the given id.
func (u *ServiceImp) HasUser(m *Model) (bool, error) {
	user, err := u.Repository.Find(m.ID)
	return user != nil, err
}

// Verify returns true if user ID and password is found in the repository.
func (u *ServiceImp) Verify(m *Model) (bool, error) {
	user, err := u.Repository.Find(m.ID)
	if err != nil {
		return false, err
	}

	return user != nil && u.PwdCrypt.Match(user.Password, m.Password), nil
}
