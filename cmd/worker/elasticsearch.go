package main

import (
	"context"

	"github.com/hardeepnarang10/ingestor/pkg/elasticsearch"
)

type elasticsearchConfig struct {
	endpoints []string
}

func initElasticSearch(ctx context.Context, esc elasticsearchConfig) (elasticsearch.ElasticSearch, func(context.Context) error) {
	e, err := elasticsearch.New(ctx,
		elasticsearch.Config{
			Addrs: esc.endpoints,
		},
	)
	if err != nil {
		report(ctx, err)
	}

	return e, e.Close
}
