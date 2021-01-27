package pwdcrypt

// Service processes the password before being persisted.
type Service interface {
	// Hash must be used to secure user' passwords before being persisted.
	Hash(pwd string) string
	// Match returns true if the password matches the hashed one.
	Match(hashed string, pwd string) bool
}
