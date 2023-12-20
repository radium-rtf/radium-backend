package filestorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func open(creds *credentials.Credentials, endpoint, region string) (*minio.Client, error) {
	secure := endpoint == "storage.radium-rtf.ru"
	options := &minio.Options{
		Secure: secure,
		Region: region,
		Creds:  creds,
	}

	return minio.New(endpoint, options)
}
