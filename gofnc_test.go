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
		stdCtx             = context.Background()
		timeOutCtx, cancel = context.WithTimeout(context.Background(), timeout)
	)
	defer cancel()

	success := func(c context.Context) error {
		t.Helper()
		r.Equal(stdCtx, c)
		return nil
	}

	expectedErr := io.EOF
	notSuccess := func(c context.Context) error {
		t.Helper()
		r.Equal(stdCtx, c)
		return expectedErr
	}

	veryLongProcess := func(c context.Context) error {
		t.Helper()
		r.Equal(timeOutCtx, c)

		time.Sleep(longProcess)
		return nil
	}

	testCases := map[string]struct {
		ctx  context.Context
		cb   func(context.Context) error
		want error
	}{
		"success":           {stdCtx, success, nil},
		"not_success":       {stdCtx, notSuccess, expectedErr},
		"very_long_process": {timeOutCtx, veryLongProcess, context.DeadlineExceeded},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			err := ctxutils.Go(tc.ctx, tc.cb)
			r.Equal(tc.want, err)
		})
	}
}
