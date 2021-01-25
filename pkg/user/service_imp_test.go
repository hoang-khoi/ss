package user

import (
	"errors"
	"ss/pkg/crypt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestServiceImp_CreateNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	pwdCryptMock := crypt.NewMockService(ctrl)

	underTest := ServiceImp{
		Repository: repositoryMock,
		PwdCrypt:   pwdCryptMock,
	}

	pwdCryptMock.EXPECT().Hash("secret").Return("hashedSecret")
	repositoryMock.EXPECT().Create(&Model{
		ID:       "koi",
		Password: "hashedSecret",
	})

	_ = underTest.CreateNewUser(&Model{
		ID:       "koi",
		Password: "secret",
	})
}

func TestServiceImp_HasUser(t *testing.T) {
	testTable := make(map[*Model]bool)
	testTable[nil] = false
	testTable[&Model{}] = true

	for user, expected := range testTable {
		hasUserTestHelper(t, user, expected)
	}
}

func TestServiceImp_Verify_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	pwdCryptMock := crypt.NewMockService(ctrl)

	underTest := ServiceImp{
		Repository: repositoryMock,
		PwdCrypt:   pwdCryptMock,
	}

	repositoryMock.EXPECT().Find("koi").Return(nil, nil)

	found, err := underTest.Verify("koi", "secret")

	assert.Nil(t, err)
	assert.False(t, found)
}

func TestServiceImp_Verify_WrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	pwdCryptMock := crypt.NewMockService(ctrl)

	underTest := ServiceImp{
		Repository: repositoryMock,
		PwdCrypt:   pwdCryptMock,
	}

	repositoryMock.EXPECT().Find("koi").Return(
		&Model{
			ID:       "koi",
			Password: "hashedSecret",
		},
		nil,
	)
	pwdCryptMock.EXPECT().Match("hashedSecret", "secret").Return(true)

	found, err := underTest.Verify("koi", "secret")

	assert.Nil(t, err)
	assert.True(t, found)
}

func TestServiceImp_Verify_FindError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	pwdCryptMock := crypt.NewMockService(ctrl)

	underTest := ServiceImp{
		Repository: repositoryMock,
		PwdCrypt:   pwdCryptMock,
	}
	errFind := errors.New("dummy")
	repositoryMock.EXPECT().Find("koi").Return(nil, errFind)

	_, err := underTest.Verify("koi", "secret")
	assert.ErrorIs(t, err, errFind)
}

func hasUserTestHelper(t *testing.T, user *Model, expected bool) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := NewMockRepository(ctrl)
	pwdCryptMock := crypt.NewMockService(ctrl)

	underTest := ServiceImp{
		Repository: repositoryMock,
		PwdCrypt:   pwdCryptMock,
	}

	repositoryMock.EXPECT().Find("koi").Return(user, nil)

	found, _ := underTest.HasUser("koi")
	assert.Equal(t, expected, found)
}
