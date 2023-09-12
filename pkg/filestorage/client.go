package filestorage

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/radium-rtf/radium-backend/config"
	"io"
)

const (
	bucket = "radium-server"
)

var policy = `{
    "Version": "2012-10-17",
    "Id": "akjsdhakshfjlashdf",
    "Statement": [
        {
            "Sid": "kjahsdkajhsdkjasda",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*"
            },
            "Action": "s3:GetObject",
            "Resource": "arn:aws:s3:::radium-server/*"
        }
    ]
}`

type Storage struct {
	client *minio.Client
}

func New(cfg config.Storage) (Storage, error) {
	creds := credentials.NewStaticV4(cfg.Id, cfg.Secret, "")
	client, err := open(creds, cfg.Endpoint, cfg.Region)
	if err != nil {
		return Storage{}, err
	}

	storage := Storage{client: client}

	err = storage.makeBucket(context.Background(), bucket, minio.MakeBucketOptions{Region: cfg.Region})
	if err != nil {
		return Storage{}, err
	}

	err = storage.client.SetBucketPolicy(context.Background(), bucket, policy)
	if err != nil {
		return Storage{}, err
	}

	return storage, err
}

func (s Storage) makeBucket(ctx context.Context, bucketName string, options minio.MakeBucketOptions) error {
	err := s.client.MakeBucket(ctx, bucketName, options)
	if err == nil {
		return err
	}
	buckets, err := s.client.ListBuckets(ctx)
	if err != nil {
		return err
	}
	for _, bucket := range buckets {
		if bucket.Name == bucketName {
			return nil
		}
	}
	return err
}

func (s Storage) PutImage(ctx context.Context, reader io.Reader, objectSize int64, contentType string) (minio.UploadInfo, error) {
	opts := minio.PutObjectOptions{ContentType: contentType}
	name := uuid.New().String()
	out, err := s.client.PutObject(ctx, bucket, name, reader, objectSize, opts)
	if err != nil {
		return minio.UploadInfo{}, err
	}
	out.Location = s.client.EndpointURL().JoinPath(out.Bucket, out.Key).String()
	return out, err
}
