package helper

import (
	"context"
	"time"
)

func Retry(
	ctx context.Context,
	maxAttempts int,
	fn func() error,
) error {

	backoff := 100 * time.Millisecond

	var err error

	for attempt := 1; attempt <= maxAttempts; attempt++ {

		err = fn()
		if err == nil {
			return nil
		}

		if attempt == maxAttempts {
			break
		}

		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-time.After(backoff):
		}

		backoff *= 2
	}

	return err
}
