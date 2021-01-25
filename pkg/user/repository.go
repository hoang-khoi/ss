package user

// Repository manages basic CRUD operations for users.
type Repository interface {
	// Create creates a new user.
	Create(u *Model) error
	// Find retrieve a user based on ID, return nil if not found.
	Find(id string) (*Model, error)
}
