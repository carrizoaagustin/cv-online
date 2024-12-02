package domain

const (
	ContentTypePDF  = "application/pdf"
	ContentTypePNG  = "image/png"
	ContentTypeJPEG = "image/jpeg"
	ContentTypeGIFT = "image/gif"
)

type FileInput struct {
	File        []byte
	Filename    string
	Folders     []string
	ContentType string
}
