package code

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextTryCancel() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "name", "main goroutine")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println(ctx.Value("name"), ",cancelFunc Called")
		wg.Done()
	}(ctx)
	cancelFunc()
	wg.Wait()
}

func ContextTryDeadline() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	ctx = context.WithValue(ctx, "name", "main goroutine")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println(ctx.Value("name"), ", time expired")
		wg.Done()
	}(ctx)
	wg.Wait()
}
func ContextTryTimeOut() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctx = context.WithValue(ctx, "name", "main goroutine")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println(ctx.Value("name"), ", time expired")
		wg.Done()
	}(ctx)
	wg.Wait()
}
