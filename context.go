package core

import (
	"context"
	"time"
)

func Context(deadline time.Duration) (context.Context, context.CancelFunc) {
	return context.WithDeadline(context.Background(), time.Now().Add(deadline))
}