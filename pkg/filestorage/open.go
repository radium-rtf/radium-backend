package filestorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"time"
)

func open(creds *credentials.Credentials, endpoint, region string) (*minio.Client, error) {
	var (
		client *minio.Client
		err    error
	)

	options := &minio.Options{
		Region: region,
		Creds:  creds,
	}

	for i := 0; i < 20; i++ {
		client, err = minio.New(endpoint, options)
		if err == nil {
			break
		}

		log.Println(err)
		time.Sleep(time.Second)
	}

	return client, err
}
