package authtoken

import (
	"github.com/dgrijalva/jwt-go"
)

// ServiceJwtHs256 implement Service using JWT HS256.
type ServiceJwtHs256 struct {
	Key []byte
}

// Generate issues a string token.
func (j *ServiceJwtHs256) Generate(m *Model) (string, error) {
	claims := jwt.StandardClaims{
		Audience:  m.ID,
		ExpiresAt: m.Expiry,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, _ := token.SignedString(j.Key)

	return signedString, nil
}

// Verify checks token's authenticity and makes sure it is a correct token.
func (j *ServiceJwtHs256) Verify(tokenString string) (*Model, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})

	if err != nil || token == nil || !token.Valid {
		return nil, err
	}

	stdClaim := token.Claims.(*jwt.StandardClaims)

	model := &Model{
		ID:     stdClaim.Audience,
		Expiry: stdClaim.ExpiresAt,
	}

	return model, nil
}
