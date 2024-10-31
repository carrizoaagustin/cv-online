package storage_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/storage"
	"github.com/stretchr/testify/mock"
)

type MockClientS3 struct {
	mock.Mock
}

func (m *MockClientS3) PutObject(ctx context.Context, input *s3.PutObjectInput, opts ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	args := m.Called(ctx, input, opts)

	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

func TestUploadFile(t *testing.T) {
	cfg := config.StorageR2{
		AccountID: "AccountID",
		AccessKey: "AccessKey",
		SecretKey: "SecretKey",
		Bucket:    "Bucket",
	}

	type Given struct {
		file     []byte
		filename string
		folders  []string
	}

	type Expected struct {
		key string
	}

	testCases := map[string]struct {
		given    Given
		expected Expected
	}{
		"File is nil":           {},
		"Folders is nil":        {},
		"Upload file to folder": {},
		"filename is empty":     {},
		"r2 error":              {},
	}

	mock := new(MockClientS3)

	fileStorage := storage.NewFileStorageServiceR2(cfg, mock)

	_, err = fileStorage.UploadFile(bytes, "carpeta/test2.png", nil)

}
