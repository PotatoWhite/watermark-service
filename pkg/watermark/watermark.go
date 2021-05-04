package watermark

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/potatowhite/watermark-service/internal"
)

type watermarkService struct{}

func (w *watermarkService) Get(_ context.Context, filter ...internal.Filter) ([]internal.Filter, error) {
	doc := internal.Document{
		Content: "book",
		Title:   "Harry Potter and Half Blood Prince",
		Author:  "J.K. Rawling",
		Topic:   "Fiction and Magic",
	}

	return []internal.Document{doc}, nil
}

func (w *watermarkService) Status(_ context.Context, ticketId string) (internal.Status, error) {
	return internal.InProgress, nil
}

func (w *watermarkService) Watermark(_ context.Context, ticketId, mark string) (int, error) {
	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(_ context.Context, doc *internal.Document) (string, error) {
	newTicketId := shortuuid.New()
	return newTicketId, nil
}

func (w *watermarkService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("Checking the Service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
