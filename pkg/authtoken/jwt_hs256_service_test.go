package authtoken

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestJwtHs256Service_Generate_Verify_CorrectToken(t *testing.T) {
	service := JwtHs256Service{Key: []byte("secret")}
	origin := Model{
		ID:     "khoi",
		Expiry: time.Now().Add(time.Hour).Unix(),
	}

	stringToken := service.Generate(&origin)
	parsed, _ := service.Verify(stringToken)

	assert.Equal(t, parsed, &origin)
}

// Note: exp is covered by jwt.StandardClaims, no need to test
func TestJwtHs256Service_Verify_WrongSignature(t *testing.T) {
	service := JwtHs256Service{Key: []byte("secret")}
	stringToken := makeInvalidSignatureToken()
	parsed, err := service.Verify(stringToken)

	assert.Nil(t, parsed)
	assert.ErrorAs(t, err, &jwt.ErrSignatureInvalid)
}

func makeInvalidSignatureToken() string {
	service := JwtHs256Service{Key: []byte("wrongSecret")}
	return service.Generate(&Model{
		ID:     "",
		Expiry: 0,
	})
}
