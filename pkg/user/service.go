package user

// Service provides necessary methods for working with users.
type Service interface {
	CreateNewUser(m *Model) error
	HasUser(id string) (bool, error)
	Verify(id string, pwd string) (bool, error)
}
