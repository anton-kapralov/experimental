package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func loopProduce(ctx context.Context, ch chan<- int) {
	t := time.NewTicker(time.Second)
	counter := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			counter++
			ch <- counter
			if counter > 4 {
				close(ch)
				return
			}
		}
	}
}

func loopConsume(ctx context.Context, name string, ch <-chan int) error {
	for {
		select {
		case <-ctx.Done():
			log.Printf("%s done", name)
			return nil
		case i, ok := <-ch:
			if !ok {
				return fmt.Errorf("%s channel closed", name)
			}
			log.Printf("%s received %d", name, i)
		}
		log.Printf("%s end of loop", name)
	}
}

func sleepAndFail() error {
	duration := 5 * time.Second
	time.Sleep(duration)
	return fmt.Errorf("%s elapsed", duration)
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	ch := make(chan int)
	eg.Go(func() error {
		return loopConsume(ctx, "loop-consume", ch)
	})
	eg.Go(sleepAndFail)
	go func() {
		loopProduce(ctx, ch)
	}()
	if err := eg.Wait(); err != nil {
		log.Fatalln("error:", err)
	}
}
