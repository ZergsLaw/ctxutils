package ctxutils

import (
	"context"
)

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
