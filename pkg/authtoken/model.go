package authtoken

// Model holds necessary information for authentication token.
type Model struct {
	ID     string
	expiry int64 // UNIX timestamp
}
