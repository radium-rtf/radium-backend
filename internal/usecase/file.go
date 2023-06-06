package usecase

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

type FileUseCase struct {
	storage filestorage.Storage
}

func NewFileUseCase(storage filestorage.Storage) FileUseCase {
	return FileUseCase{storage: storage}
}

func (uc FileUseCase) UploadFile(ctx context.Context, request entity.FileUploadRequest) (entity.FileDto, error) {
	contentType := request.Header.Header.Get("Content-Type")
	file, err := uc.storage.PutImage(ctx, request.File, request.Header.Size, contentType)
	if err != nil {
		return entity.FileDto{}, err
	}

	return entity.FileDto{Location: file.Location}, err
}
