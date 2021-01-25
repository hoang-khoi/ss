package user

// Service provides necessary methods for working with users.
type Service interface {
	CreateNewUser(m *Model) error
	// HasUser returns true if user ID is found int he repository.
	HasUser(m *Model) (bool, error)
	// Verify returns true if ID and password is found in the repository.
	Verify(m *Model) (bool, error)
}
