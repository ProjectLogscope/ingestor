package queue

import "github.com/nats-io/nats.go"

func (q *queue) Events() <-chan *nats.Msg {
	return q.eventChan
}
