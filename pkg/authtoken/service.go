package authtoken

// Service signs and verifies authentication tokens.
type Service interface {
	Generate(m *Model) string
	Verify(token string) (*Model, error)
}
