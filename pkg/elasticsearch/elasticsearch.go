package elasticsearch

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticSearch interface {
	Version(ctx context.Context) (*clusterInfo, error)
	Register(ctx context.Context, logIndex string, df DynamicFields) error
	Client() (typedClient *elasticsearch.TypedClient)
	Close(ctx context.Context) error
}
