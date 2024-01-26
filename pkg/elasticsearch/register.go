package elasticsearch

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func elasticString(s string) *string {
	return &s
}

func (c *client) Register(ctx context.Context, logIndex string, df DynamicFields) error {
	analyzer := make(map[string]types.Analyzer, len(df.keys))
	dynamicTemplates := make([]map[string]types.DynamicTemplate, len(df.keys))
	for i, key := range df.keys {
		field := key + df.suffix
		analyzer[field] = types.CustomAnalyzer{
			Type:      "custom",
			Tokenizer: "keyword",
		}

		elasticField := elasticString(field)
		dynamicTemplates[i] = map[string]types.DynamicTemplate{
			field: {
				Match: elasticField,
				Mapping: types.TextProperty{
					Type:     "text",
					Analyzer: elasticField,
				},
			},
		}
	}

	exist, err := c.typedClient.Indices.Exists(logIndex).Do(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "unable to check if log index exists", slog.Any("error", err))
		return fmt.Errorf("unable to check if log index exists: %w", err)
	}

	if !exist {
		if _, err := c.typedClient.Indices.
			Create(logIndex).
			Settings(
				&types.IndexSettings{
					Analysis: &types.IndexSettingsAnalysis{
						Analyzer: analyzer,
					},
				},
			).
			Mappings(
				&types.TypeMapping{
					DynamicTemplates: dynamicTemplates,
				},
			).Do(ctx); err != nil {
			slog.ErrorContext(ctx, "unable to create log index", slog.String("index", logIndex), slog.Any("error", err))
			return fmt.Errorf("unable to create log index %q: %w", logIndex, err)
		}
	}
	return nil
}
