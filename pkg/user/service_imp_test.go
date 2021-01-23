package user

import (
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

	_ = underTest.CreateNewUser("koi", "secret")
}

func TestServiceImp_HasUser(t *testing.T) {
	testTable := make(map[*Model]bool)
	testTable[nil] = false
	testTable[&Model{}] = true

	for user, expected := range testTable {
		hasUserTestHelper(t, user, expected)
	}
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
