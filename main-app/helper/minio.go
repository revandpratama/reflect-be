package helper

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/revandpratama/reflect/config"
)

const (
	MINIO_POST_BUCKET = "post"
	MINIO_PROFILE_BUCKET = "profile"
)

func ToRelativePath(path string) string {
	minio_endpoint := fmt.Sprintf("http://%s", config.ENV.Minio_Endpoint)

	return strings.TrimPrefix(path, minio_endpoint)
}

func CreateBucket(ctx context.Context, client *minio.Client, bucketName string) error {
	bucketExists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		NewLog().Error(fmt.Sprintf("Failed to check if bucket exists: %v", err)).ToKafka()
		return err
	}

	if !bucketExists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
		NewLog().Info(fmt.Sprintf("Bucket %s created successfully", bucketName)).ToKafka()
	} else {
		NewLog().Info(fmt.Sprintf("Bucket %s already exists", bucketName)).ToKafka()
	}

	return nil
}

func UploadObject(ctx context.Context, minioClient *minio.Client, bucketName string, rawFile *multipart.FileHeader) (string, error) {
	file, err := rawFile.Open()
	if err != nil {
		NewLog().Error(fmt.Sprintf("failed to open file: %v", err)).ToKafka()
		return "", err
	}
	defer file.Close()
	objectName := rawFile.Filename
	ext := filepath.Ext(objectName)
	baseName := strings.TrimSuffix(objectName, ext)

	// Sanitize the base name
	baseName = sanitizeFilename(baseName)

	// Add timestamp to avoid conflicts
	timestamp := time.Now().Format("20060102_150405")
	objectName = fmt.Sprintf("%s_%s%s", baseName, timestamp, ext)

	_, err = minioClient.PutObject(ctx, bucketName, objectName, file, rawFile.Size, minio.PutObjectOptions{
		ContentType: rawFile.Header.Get("Content-Type"),
	})
	if err != nil {
		NewLog().Error(fmt.Sprintf("failed to upload file: %v", err)).ToKafka()
		return "", err
	}

	url := fmt.Sprintf("http://%s/%s/%s", config.ENV.Minio_Endpoint, bucketName, objectName)
	NewLog().Info(fmt.Sprintf("File uploaded successfully: %s", url)).ToKafka()
	return url, nil
}

func DeleteObject(ctx context.Context, minioClient *minio.Client, bucketName string, urlImage string) error {
	minioEndpoint := fmt.Sprintf("http://%s/", config.ENV.Minio_Endpoint)
	objectPath := strings.TrimPrefix(urlImage, minioEndpoint)

	if !strings.HasPrefix(objectPath, bucketName+"/") {
		NewLog().Error(fmt.Sprintf("Invalid URL, does not belong to bucket: %s", urlImage)).ToKafka()
		return fmt.Errorf("invalid URL: does not belong to bucket %s", bucketName)
	}

	// Extract the objectName by removing the bucket name prefix
	objectName := strings.TrimPrefix(objectPath, bucketName+"/")

	err := minioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		NewLog().Error(fmt.Sprintf("failed to delete file: %v", err)).ToKafka()
		return err
	}

	NewLog().Info(fmt.Sprintf("File deleted successfully: %s", objectName)).ToKafka()
	return nil
}

func sanitizeFilename(name string) string {
	name = strings.TrimSpace(name)
	// Replace spaces with underscores
	name = strings.ReplaceAll(name, " ", "_")
	// Remove special characters that might cause issues
	name = regexp.MustCompile(`[<>:"/\\|?*]+`).ReplaceAllString(name, "")
	return strings.ToLower(name)
}
