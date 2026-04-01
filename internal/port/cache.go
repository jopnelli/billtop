package port

import (
	"context"
	"time"
)

// CacheStore provides local caching for query results and API responses.
type CacheStore interface {
	// Get retrieves a cached value by key. Returns nil if not found or expired.
	Get(ctx context.Context, key string) ([]byte, error)

	// Set stores a value with the given TTL.
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error

	// Close releases any resources held by the cache.
	Close() error
}
