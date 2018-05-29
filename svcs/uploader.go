package svcs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Uploader file uploader for s3
type Uploader struct {
	s3 *s3manager.Uploader
}

// NewUploader create uploader reference
func NewUploader(s *session.Session) *Uploader {
	return &Uploader{s3: s3manager.NewUploader(s)}
}

// Upload upload file to s3
func (u *Uploader) Upload(command *s3manager.UploadInput) (*s3manager.UploadOutput, error) {
	return u.s3.Upload(command)
}
