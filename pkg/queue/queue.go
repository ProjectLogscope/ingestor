package queue

import (
	"context"

	"github.com/nats-io/nats.go"
)

type Queue interface {
	Subscribe(ctx context.Context, subject string, group string) error
	Events() (msgChannel <-chan *nats.Msg)
	Close() error
}
