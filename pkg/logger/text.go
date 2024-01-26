package logger

import (
	"context"
	"log/slog"
)

type TextHandler struct {
	slog.TextHandler
}

func (h *TextHandler) Handle(ctx context.Context, r slog.Record) error {
	if meta, ok := requestInfoFromContext(ctx); ok {
		r.AddAttrs(meta)
	}
	return h.TextHandler.Handle(ctx, r)
}
