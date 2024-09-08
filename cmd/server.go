package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/junpeng.ong/protobuf-playground/api"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	// When the program is interrupted, signal the context that it should shutdown
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Start a server
	srv := api.NewService()
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("127.0.0.1", "8090"),
		Handler: srv,
	}

	// Start a goroutine that runs the server and listens to a specific port
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
