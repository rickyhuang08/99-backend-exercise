package helpers

type ContextKey string

const (
	// DefaultPageSize is the default number of items per page for pagination
	DefaultPageSize = 10
	// DefaultPage is the default page number for pagination
	DefaultPage = 1
	// DefaultTimeFormat is the default format for time representation

	//ErrorResponse is the error response message in usecase
	ErrInvalidUserPayload = "invalid user payload, name, email and password are required"
	UserNotFoundError = "user not found"
	InvalidCredentialsError = "invalid email or password"
	FailedToGenerateTokenError = "failed to generate token, please try again later"

	ErrInvalidListingPayload = "invalid listing payload, user_id and price are required"
	
	RateLimit ContextKey = "rate_limit"
)