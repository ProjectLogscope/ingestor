package queue

import (
	"fmt"

	"github.com/hardeepnarang10/ingestor/common/service"
	"github.com/nats-io/nats.go"
)

type queue struct {
	conn *nats.Conn
	subs *nats.Subscription

	eventChan chan *nats.Msg
}

func New(endpoint string) (Queue, error) {
	conn, err := nats.Connect(endpoint, nats.Name(service.GetName()))
	if err != nil {
		return nil, fmt.Errorf("unable to connect to NATS server: %w", err)
	}
	return &queue{
		conn: conn,
	}, nil
}

func (q *queue) Close() error {
	if err := q.subs.Drain(); err != nil {
		return fmt.Errorf("unable to drain NATS subscription: %w", err)
	}
	close(q.eventChan)
	return nil
}
