package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"mime/multipart"

	"github.com/radium-rtf/radium-backend/pkg/filestorage"
)

type FileUseCase struct {
	storage filestorage.Storage
	file    postgres.File
}

func NewFileUseCase(storage filestorage.Storage, file postgres.File) FileUseCase {
	return FileUseCase{storage: storage, file: file}
}

func (uc FileUseCase) UploadFile(ctx context.Context, file multipart.File, header *multipart.FileHeader) (*model.File, error) {
	contentType := header.Header.Get("Content-Type")
	info, err := uc.storage.PutImage(ctx, file, header.Size, contentType)
	if err != nil {
		return nil, err
	}
	fileEntity := &entity.File{Url: info.Location, Name: header.Filename, Type: contentType, Size: header.Size}
	err = uc.file.Create(ctx, fileEntity)
	if err != nil {
		return nil, err
	}
	return model.NewFile(fileEntity), err
}
