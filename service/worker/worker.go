package worker

import "context"

type Worker interface {
	Handler(ctx context.Context, data []byte) error
}
