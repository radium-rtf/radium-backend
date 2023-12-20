package filestorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/radium-rtf/radium-backend/config"
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

func Open(cfg config.Storage) (*minio.Client, error) {
	creds := credentials.NewStaticV4(cfg.Id, cfg.Secret, "")
	client, err := open(creds, cfg.PrivateEndpoint, cfg.Region)
	if err != nil {
		return client, err
	}
	return client, err
}
