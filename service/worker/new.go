package worker

import (
	"context"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/hardeepnarang10/ingestor/service/db/elasticsearch/store"
	"github.com/hardeepnarang10/ingestor/service/internal/msgvalidator"
	"github.com/hardeepnarang10/ingestor/service/registry"
)

var (
	w    Worker
	once sync.Once
)

type worker struct {
	workerCount      int
	enableValidation bool
	enableMocking    bool

	mv  *msgvalidator.MessageValidator
	ess store.Store
}

func New(ctx context.Context, svc *registry.ServiceRegistry) Worker {
	var err error
	once.Do(func() {
		var ess store.Store
		ess, err = store.New(ctx, svc.ElasticsearchClient, svc.ServiceConfig.StoreIndex)
		if err != nil {
			return
		}

		w = &worker{
			workerCount:      svc.ServiceConfig.WorkerCount,
			enableValidation: svc.ServiceConfig.EnableValidation,
			enableMocking:    svc.ServiceConfig.EnableMocking,

			mv:  msgvalidator.New(validator.New(validator.WithRequiredStructEnabled())),
			ess: ess,
		}
	})
	return w
}
