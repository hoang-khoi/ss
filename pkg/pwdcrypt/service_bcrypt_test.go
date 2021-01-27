package pwdcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAndMatch(t *testing.T) {
	underTest := ServiceBCrypt{}
	original := "ThisIsTopSecret1234"
	hashed := underTest.Hash(original)

	assert.NotEqual(t, original, hashed)
	assert.True(t, underTest.Match(hashed, original))
}
