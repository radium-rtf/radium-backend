package entity

import "mime/multipart"

type FileUploadRequest struct {
	File   multipart.File
	Header *multipart.FileHeader
}

type FileDto struct {
	Location string
}
