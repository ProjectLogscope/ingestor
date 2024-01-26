package document

import (
	"log/slog"
	"time"
)

type Extended struct {
	LevelTerm                      string    `json:"level_term"`
	LevelPhrase                    string    `json:"level_phrase"`
	MessageTerm                    string    `json:"message_term"`
	MessagePhrase                  string    `json:"message_phrase"`
	ResourceIDTerm                 string    `json:"resourceId_term"`
	ResourceIDPhrase               string    `json:"resourceId_phrase"`
	Timestamp                      time.Time `json:"timestamp"`
	TraceIDTerm                    string    `json:"traceId_term"`
	TraceIDPhrase                  string    `json:"traceId_phrase"`
	SpanIDTerm                     string    `json:"spanId_term"`
	SpanIDPhrase                   string    `json:"spanId_phrase"`
	CommitTerm                     string    `json:"commit_term"`
	CommitPhrase                   string    `json:"commit_phrase"`
	MetadataParentResourceIDTerm   *string   `json:"metadata_parentResourceId_term"`
	MetadataParentResourceIDPhrase *string   `json:"metadata_parentResourceId_phrase"`
}

func (e *Extended) Reduce() *Plain {
	plain := Plain{
		Level:      e.LevelPhrase,
		Message:    e.MessagePhrase,
		ResourceID: e.ResourceIDPhrase,
		Timestamp:  e.Timestamp,
		TraceID:    e.TraceIDPhrase,
		SpanID:     e.SpanIDPhrase,
		Commit:     e.CommitPhrase,
	}
	if e.MetadataParentResourceIDPhrase != nil {
		plain.Metadata = &Metadata{
			ParentResourceID: e.MetadataParentResourceIDPhrase,
		}
	}
	return &plain
}

func (e *Extended) LogValue() slog.Value {
	if e == nil {
		return slog.Value{}
	}
	attrs := []slog.Attr{
		slog.String("LevelTerm", e.LevelTerm),
		slog.String("LevelPhrase", e.LevelPhrase),
		slog.String("MessageTerm", e.MessageTerm),
		slog.String("MessagePhrase", e.MessagePhrase),
		slog.String("ResourceIDTerm", e.ResourceIDTerm),
		slog.String("ResourceIDPhrase", e.ResourceIDPhrase),
		slog.Time("TimestampPhrase", e.Timestamp),
		slog.String("TraceIDTerm", e.TraceIDTerm),
		slog.String("TraceIDPhrase", e.TraceIDPhrase),
		slog.String("SpanIDTerm", e.SpanIDTerm),
		slog.String("SpanIDPhrase", e.SpanIDPhrase),
		slog.String("CommitTerm", e.CommitTerm),
		slog.String("CommitPhrase", e.CommitPhrase),
	}
	if e.MetadataParentResourceIDTerm != nil {
		attrs = append(attrs, slog.String("MetadataParentResourceID", *e.MetadataParentResourceIDTerm))
	}
	if e.MetadataParentResourceIDPhrase != nil {
		attrs = append(attrs, slog.String("MetadataParentResourceIDPhrase", *e.MetadataParentResourceIDPhrase))
	}
	return slog.GroupValue(attrs...)
}
