package authtoken

// Model holds necessary information for authentication token.
type Model struct {
	ID     string
	Expiry int64 // UNIX timestamp
}
