package store

import (
	"context"

	"github.com/hardeepnarang10/ingestor/service/common/document"
)

type Store interface {
	Insert(ctx context.Context, documentPlain document.Plain) error
}
