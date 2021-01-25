package auth

import (
	"errors"
	"ss/pkg/authtoken"
	"ss/pkg/user"
)

// ErrWrongCredential is returned when user's id/password is wrong.
var ErrWrongCredential = errors.New("wrong user credential")

// ServiceImp implements Service with AuthTokenService and UserService.
type ServiceImp struct {
	AuthTokenService authtoken.Service
	UserService      user.Service
}

// SignIn grants authentic users tokens. If the login credential is wrong, ErrWrongCredential will be returned.
func (s *ServiceImp) SignIn(m *Model, sessionExpiryUnix int64) (string, error) {
	authentic, err := s.UserService.Verify(&user.Model{
		ID:       m.ID,
		Password: m.Password,
	})

	if err != nil {
		return "", err
	}

	if !authentic {
		return "", ErrWrongCredential
	}

	return issueNewToken(m.ID, sessionExpiryUnix, s)

}

// Refresh receives a valid token and issue another one, usually used for expiration extension.
func (s *ServiceImp) Refresh(token string, sessionExpiryUnix int64) (string, error) {
	authToken, err := s.AuthTokenService.Verify(token)
	if err != nil {
		return "", err
	}

	return issueNewToken(authToken.ID, sessionExpiryUnix, s)
}

func issueNewToken(id string, sessionExpiryUnix int64, s *ServiceImp) (string, error) {
	token, err := s.AuthTokenService.Generate(
		&authtoken.Model{
			ID:     id,
			Expiry: sessionExpiryUnix,
		},
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
