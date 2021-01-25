package authtoken

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestJwtHs256Service_Generate_Verify_CorrectToken(t *testing.T) {
	service := ServiceJwtHs256{Key: []byte("secret")}
	origin := Model{
		ID:     "khoi",
		Expiry: time.Now().Add(time.Hour).Unix(),
	}

	stringToken, _ := service.Generate(&origin)
	parsed, _ := service.Verify(stringToken)

	assert.Equal(t, parsed, &origin)
}

// Note: exp is covered by jwt.StandardClaims, no need to test
func TestJwtHs256Service_Verify_WrongSignature(t *testing.T) {
	service := ServiceJwtHs256{Key: []byte("secret")}
	stringToken := makeInvalidSignatureToken()
	parsed, err := service.Verify(stringToken)

	assert.Nil(t, parsed)
	assert.ErrorAs(t, err, &jwt.ErrSignatureInvalid)
}

func makeInvalidSignatureToken() string {
	service := ServiceJwtHs256{Key: []byte("wrongSecret")}
	token, _ := service.Generate(&Model{
		ID:     "",
		Expiry: 0,
	})

	return token
}
