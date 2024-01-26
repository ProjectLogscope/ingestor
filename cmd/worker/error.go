package main

import (
	"context"
	"log/slog"
	"os"
)

func report(ctx context.Context, err error) {
	slog.ErrorContext(ctx, "main: unrecoverable error during service initialization sequence", slog.Any("error", err))
	os.Exit(2)
}
