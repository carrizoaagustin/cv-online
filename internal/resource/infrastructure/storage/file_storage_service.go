package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"

	configAws "github.com/aws/aws-sdk-go-v2/config"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/internal/resource/domain"
)

const ErrorMessageInvalidFile string = "invalid file"
const ErrorMessageInvalidFilename string = "invalid filename"

type S3Client interface {
	PutObject(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type FileStorageServiceR2 struct {
	client S3Client
	config config.StorageR2
}

func NewS3Client(r2Config config.StorageR2) S3Client { // coverage-ignore
	cfg, err := configAws.LoadDefaultConfig(context.TODO(),
		configAws.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(r2Config.AccessKey, r2Config.SecretKey, ""),
		),
		configAws.WithRegion("auto"),
	)

	if err != nil {
		log.Fatal(err)
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2Config.AccountID))
	})
}

func NewFileStorageServiceR2(config config.StorageR2, client S3Client) domain.FileStorageService {
	return &FileStorageServiceR2{
		client: client,
		config: config,
	}
}

func (fs *FileStorageServiceR2) UploadFile(fileInput domain.FileInput) (string, error) {
	if len(fileInput.File) == 0 {
		return "", errors.New(ErrorMessageInvalidFile)
	}

	if fileInput.Filename == "" {
		return "", errors.New(ErrorMessageInvalidFilename)
	}

	var key string

	if fileInput.Folders != nil {
		key = strings.Join(fileInput.Folders, "/") + "/" + fileInput.Filename
	} else {
		key = fileInput.Filename
	}

	input := &s3.PutObjectInput{
		Bucket:      aws.String(fs.config.Bucket),
		Key:         aws.String(fileInput.Filename),
		Body:        bytes.NewReader(fileInput.File),
		ContentType: aws.String("image/png"), // Set the content type to image/jpeg (change as needed)
	}

	_, err := fs.client.PutObject(context.TODO(), input)

	if err != nil {
		return "", err
	}

	return key, nil
}
