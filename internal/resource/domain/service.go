package domain

type ResourceService interface {
	Create(resource Resource) error
}

type FileStorageService interface {
	UploadFile(file []byte, filename string, folders []string) (string, error)
}
