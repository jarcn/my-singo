package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isConcelled1(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelContext(t *testing.T) {
	ctx, close := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isConcelled1(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, ctx)
	}
	close() // context.WithCancel 返回的方法,上下文中
	time.Sleep(time.Second * 1)
}
