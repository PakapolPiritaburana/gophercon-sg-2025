package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/group"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	// hook SIGTERM, SIGINT to gracefully shutdown the service.
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt)

	g := group.New(group.WithContext(ctx))
	g.Add(newHTTPService().Run)
	g.Add(newRPCService().Run)
	g.Add(newMetricsService().Run)
	return g.Wait()
}

func newHTTPService() *Service    { return newService() }
func newRPCService() *Service     { return newService() }
func newMetricsService() *Service { return newService() }

type Service struct{}

func newService() *Service { return &Service{} }

func (svc *Service) Run(ctx context.Context) error {
	if err := svc.start(ctx); err != nil {
		return fmt.Errorf("failed to start: %w", err)
	}
	// run returns when context is done or on error
	return svc.run(ctx)
}

func (svc *Service) start(ctx context.Context) error {
	// initalise
	return nil
}

func (svc *Service) run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// do work
			time.Sleep(100 * time.Millisecond)
		}
	}
}
