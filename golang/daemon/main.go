package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

func run(ctx context.Context, wg *sync.WaitGroup) {
	t := time.NewTicker(1 * time.Second)
	startTime := time.Now()
	for {
		select {
		case <-ctx.Done():
			log.Println("Daemon is done")
			wg.Done()
			return
		case <-t.C:
			elapsed := time.Now().Sub(startTime)
			log.Println(elapsed)
		}
	}
}

func onInterrupt(cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		for s := range ch {
			os.Stderr.WriteString(fmt.Sprintln(s.String()))
			cancel()
		}
	}()
}

func startDaemon() *sync.WaitGroup {
	ctx, cancel := context.WithCancel(context.Background())
	onInterrupt(cancel)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go run(ctx, wg)

	return wg
}

type httpHandler struct{}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func shutdownHTTPServer(httpServer *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func main() {
	wg := startDaemon()

	httpServer := &http.Server{Addr: ":8080", Handler: &httpHandler{}}
	go func() { log.Fatal(httpServer.ListenAndServe()) }()

	wg.Wait()
	log.Println("Exiting")
	shutdownHTTPServer(httpServer)
}
