package observability

import (
	"context"
	"crypto/rand"
	"fmt"
	"order-service/internal/domain"
)

func newRequestID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return fmt.Sprintf("%x", b)
}

func requestIdFromContext(ctx context.Context) string {
	if id, ok := ctx.Value(domain.RequestIdKey).(string); ok {
		return id
	}

	return ""
}
