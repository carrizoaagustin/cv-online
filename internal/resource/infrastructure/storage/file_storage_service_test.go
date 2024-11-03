package storage_test

import (
	"context"
	"crypto/rand"
	"errors"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/carrizoaagustin/cv-online/config"
	"github.com/carrizoaagustin/cv-online/internal/resource/infrastructure/storage"
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
		file      []byte
		filename  string
		folders   []string
		mockValue error
	}

	type Expected struct {
		key string
		err error
	}

	filename := "testfile.png"
	folders := []string{"folder"}

	size := 16
	randomBytes := make([]byte, size)
	if _, err := rand.Read(randomBytes); err != nil {
		t.Fatal("Invalid byte generation")
	}

	errorR2Message := "Error R2"

	testCases := map[string]struct {
		given    Given
		expected Expected
	}{
		"File is nil": {
			given: Given{
				file:      nil,
				filename:  filename,
				folders:   folders,
				mockValue: nil,
			},
			expected: Expected{
				key: "",
				err: errors.New(storage.ErrorMessageInvalidFile),
			},
		},
		"Folders is nil": {
			given: Given{
				file:      randomBytes,
				filename:  filename,
				folders:   nil,
				mockValue: nil,
			},
			expected: Expected{
				key: filename,
				err: nil,
			},
		},
		"Upload file to folder": {
			given: Given{
				file:      randomBytes,
				filename:  filename,
				folders:   folders,
				mockValue: nil,
			},
			expected: Expected{
				key: strings.Join(folders, "/") + "/" + filename,
				err: nil,
			},
		},
		"Filename is empty": {
			given: Given{
				file:      randomBytes,
				filename:  "",
				folders:   folders,
				mockValue: nil,
			},
			expected: Expected{
				key: "",
				err: errors.New(storage.ErrorMessageInvalidFilename),
			},
		},
		"R2 error": {
			given: Given{
				file:      randomBytes,
				filename:  filename,
				folders:   folders,
				mockValue: errors.New(errorR2Message),
			},
			expected: Expected{
				key: "",
				err: errors.New(errorR2Message),
			},
		},
	}

	for name, caseData := range testCases {
		t.Run(name, func(t *testing.T) {
			mockS3 := new(MockClientS3)
			fileStorage := storage.NewFileStorageServiceR2(cfg, mockS3)

			mockS3.
				On("PutObject", mock.Anything, mock.Anything, mock.Anything).
				Return(&s3.PutObjectOutput{}, caseData.given.mockValue)

			key, err := fileStorage.UploadFile(caseData.given.file, caseData.given.filename, caseData.given.folders)

			require.Equal(t, caseData.expected.key, key, "Keys don't match")

			if caseData.expected.err != nil {
				require.EqualError(t, err, caseData.expected.err.Error(), "Error don't match")
			}
		})
	}
}
