package ctxutils_test

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/ZergsLaw/ctxutils"
	"github.com/stretchr/testify/require"
)

const (
	timeout     = time.Second / 2
	longProcess = time.Second * 2
)

func TestGo(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var (
		stdCtx                    = context.Background()
		timeOutCtx, timeoutCancel = context.WithTimeout(context.Background(), timeout)
		canceledCtx, cancelCtx    = context.WithCancel(context.Background())
	)
	cancelCtx()
	defer timeoutCancel()

	success := func() error {
		t.Helper()
		return nil
	}

	expectedErr := io.EOF
	notSuccess := func() error {
		t.Helper()
		return expectedErr
	}

	veryLongProcess := func() error {
		t.Helper()

		time.Sleep(longProcess)
		return nil
	}

	testCases := map[string]struct {
		ctx  context.Context
		cb   func() error
		want error
	}{
		"success":             {stdCtx, success, nil},
		"not_success":         {stdCtx, notSuccess, expectedErr},
		"very_long_process":   {timeOutCtx, veryLongProcess, context.DeadlineExceeded},
		"context_is_canceled": {canceledCtx, notSuccess, context.Canceled},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			err := ctxutils.Go(tc.ctx, tc.cb)
			r.Equal(tc.want, err)
		})
	}
}
