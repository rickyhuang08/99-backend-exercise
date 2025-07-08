package auth

// Context key for request metadata
type ContextKey string

const (
	RequestIDKey ContextKey = "request_id"
	UserKey ContextKey = "user"
)