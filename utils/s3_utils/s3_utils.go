package s3_utils

import (
	"io"
	"os"

	aws_aws "github.com/aws/aws-sdk-go/aws"
	aws_credentials "github.com/aws/aws-sdk-go/aws/credentials"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_s3manager "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cheggaaa/pb/v3"
)

func Upload(
	access_key string,
	secret_key string,
	region string,
	endpoint string,
	bucket_name string,
	key string,
	f io.ReadSeeker,
) error {
	awsConfig := aws_aws.Config{
		Credentials: aws_credentials.NewStaticCredentials(
			access_key,
			secret_key,
			"",
		),
	}
	if region != "" {
		awsConfig.Region = aws_aws.String(region)
	}
	if endpoint != "" {
		awsConfig.Region = aws_aws.String(string("us-east-1"))
		awsConfig.S3ForcePathStyle = aws_aws.Bool(true)
		awsConfig.Endpoint = aws_aws.String(endpoint)
	}
	session, err := aws_session.NewSession(
		&awsConfig,
	)
	if err != nil {
		return err
	}
	uploader := aws_s3manager.NewUploader(session, func(u *aws_s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = 10            // default is 5
	})

	size, _ := f.Seek(0, os.SEEK_END)
	f.Seek(0, 0)

	bar := pb.Full.Start64(size)

	// create proxy reader
	barReader := bar.NewProxyReader(f)

	_, err = uploader.Upload(&aws_s3manager.UploadInput{
		Bucket: aws_aws.String(bucket_name),
		ACL:    aws_aws.String("private"),
		Key:    aws_aws.String(key),
		Body:   barReader,
	})
	if err != nil {
		return err
	}

	bar.Finish()

	return nil
}
