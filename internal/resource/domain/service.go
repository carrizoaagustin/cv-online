package domain

type ResourceService interface {
	Create(resource Resource) error
}

type FileStorageService interface {
	UploadFile(file FileInput) (string, error)
}
