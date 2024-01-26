package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/hardeepnarang10/ingestor/cmd/worker/internal/config"
	"github.com/hardeepnarang10/ingestor/common/service"
	"github.com/hardeepnarang10/ingestor/service/registry"
)

func main() {
	cfg := config.LoadConfig()
	service.SetName(cfg.ServiceName)
	initLogger(cfg.ServiceLogFilepath, cfg.ServiceLogLevel.GetLevel(), cfg.ServiceLogAddSource)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	var queueDrain func(context.Context)
	var elasticsearchClose func(context.Context) error
	svc := &registry.ServiceRegistry{}
	svc.QueueClient, queueDrain = initQueue(ctx, queueConfig{
		endpoint: cfg.NATSServerEndpoint,
		subject:  cfg.NATSSubjectIngest,
		group:    cfg.NATSGroupIngest,
	})
	svc.ElasticsearchClient, elasticsearchClose = initElasticSearch(
		ctx, elasticsearchConfig{
			endpoints: cfg.ElasticsearchEndpoints,
		})
	svc.ServiceConfig = registry.ServiceConfig{
		StoreIndex:       cfg.ElasticsearchIndex,
		WorkerCount:      cfg.WorkerMaxCount,
		EnableValidation: cfg.ServiceMessageValidate,
		EnableMocking:    cfg.ServiceEnableMocking,
	}

	startWorker := initWorker(ctx, svc, cfg.WorkerTimeout)

	defer func() {
		graceCtx, cancel := context.WithTimeout(context.Background(), cfg.ServiceGracePeriod)
		defer cancel()
		queueDrain(graceCtx)
		if err := elasticsearchClose(graceCtx); err != nil {
			slog.ErrorContext(graceCtx, "error closing elasticsearch client", slog.Any("error", err))
			report(graceCtx, err)
		}
	}()

	startWorker()
}
