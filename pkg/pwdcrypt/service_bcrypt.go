package pwdcrypt

import "golang.org/x/crypto/bcrypt"

// ServiceBCrypt implements Service using Go's bcrypt package.
type ServiceBCrypt struct{}

// Hash secures the password with salt
func (ServiceBCrypt) Hash(pwd string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hashed)
}

// Match returns true if the password matches the hashed one.
func (ServiceBCrypt) Match(hashed string, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)) == nil
}
