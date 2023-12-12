package filestorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func open(creds *credentials.Credentials, endpoint, region string) (*minio.Client, error) {
	options := &minio.Options{
		Region: region,
		Creds:  creds,
	}

	return minio.New(endpoint, options)
}
