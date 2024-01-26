package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/hardeepnarang10/ingestor/service/registry"
	"github.com/hardeepnarang10/ingestor/service/worker"
)

func initWorker(ctx context.Context, svc *registry.ServiceRegistry, workTimeout time.Duration) func() {
	if workTimeout == 0 {
		report(ctx, errors.New("one or more required value missing or empty"))
	}

	wc := worker.New(ctx, svc)
	eventLoop := svc.QueueClient.Events()

	return func() {
		slog.InfoContext(ctx,
			"starting ingestion worker",
			slog.Bool("operation.mock", svc.ServiceConfig.EnableMocking),
			slog.Int("workers.max", svc.ServiceConfig.WorkerCount))

		var wg = sync.WaitGroup{}
		defer wg.Wait()

		semaphore := make(chan struct{}, svc.ServiceConfig.WorkerCount)
		for {
			select {
			case <-ctx.Done():
				slog.Info("shutting down worker group: main context cancelled")
				return

			case event := <-eventLoop:
				semaphore <- struct{}{}
				wg.Add(1)
				defer wg.Done()

				ctxTimeout, cancel := context.WithTimeoutCause(ctx, workTimeout, fmt.Errorf("message ingestion context expired after %q", workTimeout.String()))
				defer cancel()

				if err := wc.Handler(ctxTimeout, event.Data); err != nil {
					slog.ErrorContext(ctxTimeout, "handler returned error", slog.String("msg.data", string(event.Data)), slog.Any("error", err))
				}

				<-semaphore
			}
		}
	}
}
