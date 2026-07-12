package observability

import (
	"crypto/rand"
	"fmt"
)

func newRequestID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])
	return fmt.Sprintf("%x", b)
}
