package auth

import (
	"errors"
	"ss/pkg/authtoken"
	"ss/pkg/user"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestServiceImp_SignIn_UserVerifyError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authTokenService := authtoken.NewMockService(ctrl)
	userService := user.NewMockService(ctrl)

	errUserService := errors.New("dummy")

	underTest := ServiceImp{
		AuthTokenService: authTokenService,
		UserService:      userService,
	}

	userService.EXPECT().Verify(&user.Model{
		ID:       "koi",
		Password: "secret",
	}).Return(false, errUserService)

	_, err := underTest.SignIn(
		&Model{
			ID:       "koi",
			Password: "secret",
		},
		12345,
	)

	assert.ErrorIs(t, err, errUserService)
}

func TestServiceImp_SignIn_UserVerifyFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authTokenService := authtoken.NewMockService(ctrl)
	userService := user.NewMockService(ctrl)

	underTest := ServiceImp{
		AuthTokenService: authTokenService,
		UserService:      userService,
	}

	userService.EXPECT().Verify(&user.Model{
		ID:       "koi",
		Password: "secret",
	}).Return(false, nil)

	_, err := underTest.SignIn(
		&Model{
			ID:       "koi",
			Password: "secret",
		},
		12345,
	)

	assert.ErrorIs(t, err, ErrWrongCredential)
}

func TestServiceImp_SignIn_TokenGenerationFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authTokenService := authtoken.NewMockService(ctrl)
	userService := user.NewMockService(ctrl)

	underTest := ServiceImp{
		AuthTokenService: authTokenService,
		UserService:      userService,
	}

	errTokenGeneration := errors.New("dummy")

	userService.EXPECT().Verify(&user.Model{
		ID:       "koi",
		Password: "secret",
	}).Return(true, nil)

	authTokenService.EXPECT().Generate(&authtoken.Model{
		ID:     "koi",
		Expiry: 12345,
	}).Return("", errTokenGeneration)

	_, err := underTest.SignIn(
		&Model{
			ID:       "koi",
			Password: "secret",
		},
		12345,
	)

	assert.ErrorIs(t, err, errTokenGeneration)
}

func TestServiceImp_SignIn_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authTokenService := authtoken.NewMockService(ctrl)
	userService := user.NewMockService(ctrl)

	underTest := ServiceImp{
		AuthTokenService: authTokenService,
		UserService:      userService,
	}

	userService.EXPECT().Verify(&user.Model{
		ID:       "koi",
		Password: "secret",
	}).Return(true, nil)

	authTokenService.EXPECT().Generate(&authtoken.Model{
		ID:     "koi",
		Expiry: 12345,
	}).Return("token", nil)

	token, _ := underTest.SignIn(
		&Model{
			ID:       "koi",
			Password: "secret",
		},
		12345,
	)

	assert.Equal(t, "token", token)
}

func TestServiceImp_Refresh_InvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authTokenService := authtoken.NewMockService(ctrl)
	userService := user.NewMockService(ctrl)

	underTest := ServiceImp{
		AuthTokenService: authTokenService,
		UserService:      userService,
	}

	errTokenVerification := errors.New("dummy")
	authTokenService.EXPECT().Verify("token").Return(nil, errTokenVerification)

	_, err := underTest.Refresh("token", 12345)

	assert.ErrorIs(t, err, errTokenVerification)
}

func TestServiceImp_Refresh_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authTokenService := authtoken.NewMockService(ctrl)
	userService := user.NewMockService(ctrl)

	underTest := ServiceImp{
		AuthTokenService: authTokenService,
		UserService:      userService,
	}

	authTokenService.EXPECT().Verify("token").Return(
		&authtoken.Model{
			ID:     "koi",
			Expiry: 12345,
		},
		nil,
	)

	authTokenService.EXPECT().Generate(&authtoken.Model{
		ID:     "koi",
		Expiry: 12345,
	}).Return("newToken", nil)

	token, _ := underTest.Refresh("token", 12345)
	assert.Equal(t, "newToken", token)
}
