package s3

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/utils/file_utils"
	"github.com/sikalabs/tergum/utils/s3_utils"
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
	return s3_utils.Upload(
		t.AccessKey,
		t.SecretKey,
		t.Region,
		t.Endpoint,
		t.BucketName,
		file_utils.GetFileName(t.Prefix, t.Suffix),
		data,
	)
}
