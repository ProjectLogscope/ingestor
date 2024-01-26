package registry

import (
	"github.com/hardeepnarang10/ingestor/pkg/elasticsearch"
	"github.com/hardeepnarang10/ingestor/pkg/queue"
)

type ServiceRegistry struct {
	ServiceConfig       ServiceConfig
	QueueClient         queue.Queue
	ElasticsearchClient elasticsearch.ElasticSearch
}

type ServiceConfig struct {
	StoreIndex       string
	WorkerCount      int
	EnableValidation bool
	EnableMocking    bool
}
