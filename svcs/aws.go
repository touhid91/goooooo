package svcs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Dial connects to the provide region of aws
func Dial(region string) (*session.Session, error) {
	ses, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return nil, err
	}

	return ses, nil
}

// S3 get provider reference for S3
func S3(s *session.Session) *s3.S3 {
	return s3.New(s)
}
