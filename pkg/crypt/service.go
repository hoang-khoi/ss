package crypt

// Service processes the password before being persisted.
type Service interface {
	Hash(pwd string) string
	Match(hashed string, pwd string) bool
}
