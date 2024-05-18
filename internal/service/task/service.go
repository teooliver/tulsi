package task

import "context"

type taskRepo interface {
	Get(ctx context.Context, id string) (storage.Metadata, error)
	Create(ctx context.Context, params storage.CreateParams)
}
