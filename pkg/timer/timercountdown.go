package timer

import (
	"context"
	"time"
)

func Countdown(countDown int, ctxFunc func(ctx context.Context, cancel context.CancelFunc), countFunc func(int)) {
	if countDown < 0 {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	if ctxFunc != nil {
		ctxFunc(ctx, cancel)
	}
	for i := countDown; i >= 0; i-- {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			if countFunc != nil {
				countFunc(i)
			}
		}
	}
}
