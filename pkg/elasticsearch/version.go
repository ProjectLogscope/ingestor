package elasticsearch

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type clusterInfo struct {
	Name    string
	Version types.ElasticsearchVersionInfo
}

func (c *client) Version(ctx context.Context) (*clusterInfo, error) {
	ci, err := c.typedClient.Info().Do(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "unable to get cluster info", slog.Any("error", err))
		return nil, fmt.Errorf("unable to get cluster info: %w", err)
	}
	return &clusterInfo{
		Name:    ci.ClusterName,
		Version: ci.Version,
	}, nil
}
