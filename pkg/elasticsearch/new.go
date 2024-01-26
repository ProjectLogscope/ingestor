package elasticsearch

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
)

type client struct {
	typedClient *elasticsearch.TypedClient
}

func New(ctx context.Context, cfg Config) (ElasticSearch, error) {
	tc, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: cfg.Addrs,
		Logger: &elastictransport.JSONLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
		RetryOnStatus: []int{
			http.StatusTooManyRequests,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
		},
		DiscoverNodesOnStart:  true,
		DiscoverNodesInterval: time.Minute,
	})
	if err != nil {
		slog.ErrorContext(ctx, "unable to create typed client", slog.Any("config", cfg), slog.Any("error", err))
		return nil, fmt.Errorf("unable to create typed client with config %+v: %w", cfg, err)
	}

	return &client{
		typedClient: tc,
	}, nil
}

func (c *client) Client() *elasticsearch.TypedClient {
	return c.typedClient
}

func (c *client) Close(ctx context.Context) error {
	// TODO: Close client connection
	return nil
}
