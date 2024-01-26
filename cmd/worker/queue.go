package main

import (
	"context"

	"github.com/hardeepnarang10/ingestor/pkg/queue"
)

type queueConfig struct {
	endpoint string
	subject  string
	group    string
}

func initQueue(ctx context.Context, qc queueConfig) (queue.Queue, func(context.Context)) {
	q, err := queue.New(qc.endpoint)
	if err != nil {
		report(ctx, err)
	}

	if err := q.Subscribe(ctx, qc.subject, qc.group); err != nil {
		report(ctx, err)
	}

	return q, func(ctx context.Context) {
		if err := q.Close(); err != nil {
			report(ctx, err)
		}
	}
}
