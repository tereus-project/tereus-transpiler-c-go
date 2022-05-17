package services

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioService struct {
	bucket string

	client *minio.Client
}

func NewMinioService(endpoint string, accessKey string, secretKey string, bucket string) (*MinioService, error) {
	// Initialize minio client object.
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	// Create tereus S3 bucket if it doesn't exist
	exists, err := client.BucketExists(context.Background(), bucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		err = client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &MinioService{
		bucket: bucket,
		client: client,
	}, nil
}

var remixPrefix = "remix"

func (s *MinioService) GetFiles(id string) <-chan string {
	files := make(chan string)
	prefix := fmt.Sprintf("%s/%s/", remixPrefix, id)

	go func() {
		for object := range s.client.ListObjects(
			context.Background(),
			s.bucket,
			minio.ListObjectsOptions{
				Prefix:    prefix,
				Recursive: true,
			},
		) {
			files <- strings.TrimPrefix(object.Key, prefix)
		}

		close(files)
	}()

	return files
}

func (s *MinioService) GetFile(id string, filepath string) (string, error) {
	objectPath := fmt.Sprintf("%s/%s/%s", remixPrefix, id, filepath)

	object, err := s.client.GetObject(
		context.Background(),
		s.bucket,
		objectPath,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return "", err
	}
	defer object.Close()

	dir, err := os.MkdirTemp("", fmt.Sprintf("%s-", id))
	if err != nil {
		return "", fmt.Errorf("Failed to create temp dir: %s", err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s", dir, filepath))
	if err != nil {
		return "", fmt.Errorf("failed to create file: %s", err)
	}

	_, err = io.Copy(f, object)
	if err != nil {
		return "", fmt.Errorf("failed to copy object to '%s': %s", f.Name(), err)
	}

	return f.Name(), nil
}

func (s *MinioService) PutFile(id string, filepath string, content string) error {
	objectPath := fmt.Sprintf("%s-results/%s/%s", remixPrefix, id, filepath)

	_, err := s.client.PutObject(
		context.Background(),
		s.bucket,
		objectPath,
		strings.NewReader(content),
		int64(len(content)),
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}

	return nil
}
