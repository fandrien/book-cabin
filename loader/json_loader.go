package loader

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func LoadJSON[T any](path string) (T, error) {
	var obj T

	data, err := os.ReadFile(path)
	if err != nil {
		return obj, err
	}

	if err := json.Unmarshal(data, &obj); err != nil {
		return obj, err
	}

	return obj, nil
}

func RandomDelay(ctx context.Context, minMs, maxMs int) error {
	if minMs > maxMs {
		minMs, maxMs = maxMs, minMs
	}

	delay := time.Duration(rand.Intn(maxMs-minMs+1)+minMs) * time.Millisecond

	select {
	case <-ctx.Done():
		return ctx.Err()

	case <-time.After(delay):
		return nil
	}
}
