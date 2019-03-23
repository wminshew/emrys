package check

import (
	"context"
)

// ContextCanceled checks if the context has been canceled
func ContextCanceled(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
