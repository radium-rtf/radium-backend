package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/model"
	"mime/multipart"

	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

type FileUseCase struct {
	storage filestorage.Storage
}

func NewFileUseCase(storage filestorage.Storage) FileUseCase {
	return FileUseCase{storage: storage}
}

func (uc FileUseCase) UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader) (model.File, error) {
	contentType := header.Header.Get("Content-Type")
	info, err := uc.storage.PutImage(ctx, file, header.Size, contentType)
	if err != nil {
		return model.File{}, err
	}

	return model.File{Location: info.Location}, err
}
