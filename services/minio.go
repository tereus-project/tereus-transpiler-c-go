package services

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tereus-project/tereus-remixer-c-go/env"
)

type MinioService struct {
	client *minio.Client
}

func NewMinioService() (*MinioService, error) {
	var err error
	service := &MinioService{}

	// Initialize minio client object.
	service.client, err = minio.New(env.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(env.S3AccessKey, env.S3SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	// Create tereus S3 bucket if it doesn't exist
	exists, err := service.client.BucketExists(context.Background(), env.S3Bucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		err = service.client.MakeBucket(context.Background(), env.S3Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return service, nil
}

var remixPrefix = "remix"

func (s *MinioService) GetFiles(id string) <-chan string {
	files := make(chan string)
	prefix := fmt.Sprintf("%s/%s/", remixPrefix, id)

	go func() {
		for object := range s.client.ListObjects(
			context.Background(),
			env.S3Bucket,
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
		env.S3Bucket,
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
		env.S3Bucket,
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
