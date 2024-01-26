package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/hardeepnarang10/ingestor/service/common/document"
)

func (w *worker) Handler(ctx context.Context, data []byte) error {
	var documentPlain document.Plain
	if err := json.Unmarshal(data, &documentPlain); err != nil {
		slog.DebugContext(ctx, "unable to parse incoming request body to log message type", slog.String("data", string(data)), slog.Any("error", err))
		return fmt.Errorf("unable to parse incoming request body to log message type: %w", err)
	}

	if w.enableValidation {
		if err := documentPlain.Validate(ctx, w.mv); err != nil {
			slog.DebugContext(ctx, "message validation failure", slog.Any("error", err))
			return fmt.Errorf("message validation failure: %w", err)
		}
	}

	if w.enableMocking {
		slog.InfoContext(ctx, "Mocking Document Ingestion", slog.Any("document", documentPlain))
	} else {
		if err := w.ess.Insert(ctx, documentPlain); err != nil {
			slog.ErrorContext(ctx, "unable to insert log message into elasticsearch", slog.Any("document_plain", documentPlain), slog.Any("error", err))
			return fmt.Errorf("unable to insert log message into elasticsearch: %w", err)
		}
		slog.DebugContext(ctx, "log message inserted into elasticsearch", slog.Any("document_plain", documentPlain))
	}

	return nil
}
