package authtoken

// Service signs and verifies authentication tokens.
type Service interface {
	// Generate issues a string token.
	Generate(m *Model) (string, error)
	// Verify checks token's authenticity and makes sure it is a correct token. Returned the verified parsed token.
	Verify(token string) (*Model, error)
}
