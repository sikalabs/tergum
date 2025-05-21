package s3

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/utils/file_utils"
	"github.com/sikalabs/tergum/utils/s3_utils"
)

type S3Target struct {
	AccessKey     string `yaml:"AccessKey" json:"AccessKey,omitempty"`
	SecretKey     string `yaml:"SecretKey" json:"SecretKey,omitempty"`
	Region        string `yaml:"Region" json:"Region,omitempty"`
	Endpoint      string `yaml:"Endpoint" json:"Endpoint,omitempty"`
	BucketName    string `yaml:"BucketName" json:"BucketName,omitempty"`
	Prefix        string `yaml:"Prefix" json:"Prefix,omitempty"`
	Suffix        string `yaml:"Suffix" json:"Suffix,omitempty"`
	UploadRetries int    `yaml:"UploadRetries" json:"UploadRetries,omitempty"`
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
	if t.UploadRetries < 0 {
		return fmt.Errorf("S3Target requires Retries >= 0")
	}
	return nil
}

func (t S3Target) Save(data io.ReadSeeker) error {
	var err error
	tries := 1 + t.UploadRetries
	for i := 0; i < tries; i++ {
		err = s3_utils.Upload(
			t.AccessKey,
			t.SecretKey,
			t.Region,
			t.Endpoint,
			t.BucketName,
			file_utils.GetFileName(t.Prefix, t.Suffix),
			data,
		)
		if err == nil {
			return nil
		}
	}
	return err
}
