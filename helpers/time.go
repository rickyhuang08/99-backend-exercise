package helpers

import "time"

// TimeProvider defines an interface for time operations
type TimeProvider interface {
	Now() time.Time
	Since(t time.Time) time.Duration
}

// RealTimeProvider is the actual implementation
type RealTimeProvider struct{}

// NewRealTimeProvider creates a new instance of RealTimeProvider
func NewRealTimeProvider() TimeProvider {
	return &RealTimeProvider{}
}

// Now returns the current time
func (RealTimeProvider) Now() time.Time {
	return time.Now()
}

// Since returns the duration since a given time
func (RealTimeProvider) Since(t time.Time) time.Duration {
	return time.Since(t)
}