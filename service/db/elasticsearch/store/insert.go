package store

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/hardeepnarang10/ingestor/service/common/document"
)

func (s *store) Insert(ctx context.Context, documentPlain document.Plain) error {
	if _, err := s.client.Client().
		Index(s.logIndex).
		Id(uuid.NewString()).
		Request(documentPlain.Extend()).
		Do(ctx); err != nil {
		slog.ErrorContext(ctx,
			"unable to put message into elasticsearch",
			slog.String("index", s.logIndex),
			slog.Any("document", documentPlain),
			slog.Any("error", err),
		)
		return fmt.Errorf("unable to put message into elasticsearch: %w", err)
	}
	return nil
}
