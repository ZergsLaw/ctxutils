// Package ctxutils contains utils and helpers for working with context.
package ctxutils

import (
	"context"
)

// Go call function and monitors context if context is done, return context error
// else return result from callback.
func Go(ctx context.Context, fn func(context.Context) error) error {
	errc := make(chan error, 1)

	go func() {
		errc <- fn(ctx)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errc:
		return err
	}
}
