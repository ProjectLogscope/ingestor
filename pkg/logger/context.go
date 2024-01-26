package logger

import (
	"context"
	"log/slog"
)

type ctxKeyRequest struct{}

var ContextKeyRequest ctxKeyRequest = ctxKeyRequest{}

func requestInfoFromContext(ctx context.Context) (slog.Attr, bool) {
	v, ok := ctx.Value(ContextKeyRequest).(slog.Attr)
	return v, ok
}
