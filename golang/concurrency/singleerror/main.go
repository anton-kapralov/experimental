package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"time"
)

func loop(ctx context.Context, name string) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}
		duration := time.Duration((rand.Intn(1000) + 1000) * int(time.Millisecond))
		time.Sleep(duration)
		log.Printf("%s", name)
	}
}

func sleepAndFail() error {
	duration := 3500 * time.Millisecond
	time.Sleep(duration)
	return fmt.Errorf("%s elapsed", duration)
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 3; i++ {
		id := i
		eg.Go(func() error {
			return loop(ctx, fmt.Sprintf("loop-%d", id))
		})
	}
	eg.Go(sleepAndFail)
	if err := eg.Wait(); err != nil {
		log.Fatalln("error:", err)
	}
}
