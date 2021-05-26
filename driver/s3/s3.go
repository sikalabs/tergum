package s3

import (
	"bytes"
	"errors"

	aws_aws "github.com/aws/aws-sdk-go/aws"
	aws_credentials "github.com/aws/aws-sdk-go/aws/credentials"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sikalabs/tergum/utils/file_utils"
)

type S3 struct {
	AccessKey  string
	SecretKey  string
	Region     string
	Endpoint   string
	BucketName string
	Prefix     string
	Suffix     string
}

func ValidateS3(config S3) error {
	if config.AccessKey == "" {
		return errors.New("s3 requires accessKey")
	}
	if config.SecretKey == "" {
		return errors.New("s3 requires secretKey")
	}
	if config.Region == "" && config.Endpoint == "" {
		return errors.New("s3 requires region or endpoint")
	}
	if config.BucketName == "" {
		return errors.New("s3 requires bucketName")
	}
	if config.Prefix == "" {
		return errors.New("s3 requires prefix")
	}
	if config.Suffix == "" {
		return errors.New("s3 requires suffix")
	}
	return nil
}

func SaveS3(config S3, data []byte) error {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			config.AccessKey,
			config.SecretKey,
			"",
		),
	}
	if config.Region != "" {
		awsConfig.Region = aws_aws.String(string(config.Region))
	}
	if config.Endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(string(config.Endpoint))
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}
	uploader := aws_s3manager.NewUploader(session)
	_, err = uploader.Upload(&aws_s3manager.UploadInput{
		Bucket: aws_aws.String(config.BucketName),
		ACL:    aws_aws.String("private"),
		Key:    aws_aws.String(file_utils.GetFileName(config.Prefix, config.Suffix)),
		Body:   bytes.NewReader(data),
	})
	if err != nil {
		return err
	}
	return nil
}
