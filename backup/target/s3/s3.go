package s3

import (
	"fmt"
	"io"

	aws_aws "github.com/aws/aws-sdk-go/aws"
	aws_credentials "github.com/aws/aws-sdk-go/aws/credentials"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sikalabs/tergum/utils/file_utils"
)

type S3Target struct {
	AccessKey  string `yaml:"AccessKey"`
	SecretKey  string `yaml:"SecretKey"`
	Region     string `yaml:"Region"`
	Endpoint   string `yaml:"Endpoint"`
	BucketName string `yaml:"BucketName"`
	Prefix     string `yaml:"Prefix"`
	Suffix     string `yaml:"Suffix"`
}

func (t S3Target) Validate() error {
	if t.AccessKey == "" {
		return fmt.Errorf("S3Target need to have a AccessKey")
	}
	if t.SecretKey == "" {
		return fmt.Errorf("S3Target requires SecretKey")
	}
	if t.Region == "" && t.Endpoint == "" {
		return fmt.Errorf("S3Target requires region or Endpoint")
	}
	if t.BucketName == "" {
		return fmt.Errorf("S3Target requires BucketName")
	}
	if t.Prefix == "" {
		return fmt.Errorf("S3Target requires Prefix")
	}
	if t.Suffix == "" {
		return fmt.Errorf("S3Target requires Suffix")
	}
	return nil
}

func (t S3Target) Save(data io.ReadSeeker) error {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			t.AccessKey,
			t.SecretKey,
			"",
		),
	}
	if t.Region != "" {
		awsConfig.Region = aws_aws.String(string(t.Region))
	}
	if t.Endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(string(t.Endpoint))
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}
	uploader := aws_s3manager.NewUploader(session)
	_, err = uploader.Upload(&aws_s3manager.UploadInput{
		Bucket: aws_aws.String(t.BucketName),
		ACL:    aws_aws.String("private"),
		Key:    aws_aws.String(file_utils.GetFileName(t.Prefix, t.Suffix)),
		Body:   data,
	})
	if err != nil {
		return err
	}
	return nil
}
