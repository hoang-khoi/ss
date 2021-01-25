package auth

// Service provides methods for session-related features.
type Service interface {
	SignIn(m *Model, sessionExpiryUnix int64) (string, error)
	Refresh(token string, sessionExpiryUnix int64) (string, error)
}
