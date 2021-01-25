package authtoken

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtHs256Service implement Service using JWT HS256.
type JwtHs256Service struct {
	Key []byte
}

// Generate issues a string token.
func (j *JwtHs256Service) Generate(m *Model) string {
	claims := jwt.StandardClaims{
		Audience:  m.ID,
		ExpiresAt: m.expiry,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, _ := token.SignedString(j.Key)

	return signedString
}

// Verify checks token's authenticity and makes sure it is a correct token.
func (j *JwtHs256Service) Verify(tokenString string) (*Model, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})

	if token == nil || !token.Valid {
		return nil, err
	}

	stdClaim := token.Claims.(*jwt.StandardClaims)

	model := &Model{
		ID:     stdClaim.Audience,
		expiry: stdClaim.ExpiresAt,
	}

	return model, nil
}
