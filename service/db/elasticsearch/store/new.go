package store

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/hardeepnarang10/ingestor/pkg/elasticsearch"
)

type store struct {
	client   elasticsearch.ElasticSearch
	logIndex string
}

func New(ctx context.Context, client elasticsearch.ElasticSearch, logIndex string) (Store, error) {
	if err := client.Register(ctx, logIndex, elasticsearch.DefaultDynamicFields); err != nil {
		slog.ErrorContext(ctx, "failed to register index", slog.String("index", logIndex), slog.Any("error", err))
		return nil, fmt.Errorf("failed to register index %q: %w", logIndex, err)
	}

	return &store{
		client:   client,
		logIndex: logIndex,
	}, nil
}
