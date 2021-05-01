package watermark

import (
	"context"
	"github.com/potatowhite/watermark-service/internal"
)

type Service interface {
	Get(ctx context.Context, filter ...internal.Filter) ([]internal.Document, error)
	Status(ctx context.Context, ticketIT string) (internal.Status, error, error)
	Watermark(ctx context.Context, ticketId, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (string, error)
	ServiceStatus(ctx context.Context) (int, error, error)
}
