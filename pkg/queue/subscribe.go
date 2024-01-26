package queue

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/nats-io/nats.go"
)

func (q *queue) Subscribe(ctx context.Context, subject string, group string) error {
	msgChan := make(chan *nats.Msg)
	subs, err := q.conn.ChanQueueSubscribe(subject, group, msgChan)
	if err != nil {
		slog.ErrorContext(ctx, "unable to subscribe to NATS subject as group member", slog.String("subject", subs.Subject), slog.String("queue", subs.Queue), slog.Any("error", err))
		return fmt.Errorf("unable to subscribe to subject %q as group member %q: %w", subs.Subject, subs.Queue, err)
	}
	q.eventChan = msgChan

	return nil
}
