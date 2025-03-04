package adapter

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/revandpratama/reflect/config"
	"github.com/revandpratama/reflect/helper"
)

type MinioOption struct {
	minioClient *minio.Client
}

func (m *MinioOption) Start(a *Adapter) error {

	client, err := minio.New("", &minio.Options{
		Creds:  credentials.NewStaticV4(config.ENV.Minio_AccessKey, config.ENV.Minio_SecretKey, ""),
		Secure: config.ENV.Minio_UseSSL == "true",
	})
	if err != nil {
		return fmt.Errorf("failed to connect to minio: %w", err)
	}

	a.MinioClient = client
	m.minioClient = client

	buckets := []string{
		"profile",
		"post",
	}
	for _, bucket := range buckets {
		err := helper.CreateBucket(context.Background(), m.minioClient, bucket)
		if err != nil {
			return err
		}
	}

	helper.NewLog().Info("minio storage running").ToKafka()
	return nil
}

func (m *MinioOption) Stop() error {
	return nil
}
