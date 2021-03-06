package storage

import "context"

type StorageService interface {
	Upload(ctx context.Context, content File) (*StorageResult, error)
	Delete(ctx context.Context, fileName string) (bool, error)
}
