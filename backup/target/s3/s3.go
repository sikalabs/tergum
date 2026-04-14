package s3

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sikalabs/tergum/utils/file_utils"
	"github.com/sikalabs/tergum/utils/s3_utils"
)

type S3Retention struct {
	KeepCount int `yaml:"KeepCount" json:"KeepCount,omitempty"`
	KeepDays  int `yaml:"KeepDays" json:"KeepDays,omitempty"`
}

type S3Target struct {
	AccessKey     string       `yaml:"AccessKey" json:"AccessKey,omitempty"`
	SecretKey     string       `yaml:"SecretKey" json:"SecretKey,omitempty"`
	Region        string       `yaml:"Region" json:"Region,omitempty"`
	Endpoint      string       `yaml:"Endpoint" json:"Endpoint,omitempty"`
	BucketName    string       `yaml:"BucketName" json:"BucketName,omitempty"`
	Prefix        string       `yaml:"Prefix" json:"Prefix,omitempty"`
	Suffix        string       `yaml:"Suffix" json:"Suffix,omitempty"`
	UploadRetries int          `yaml:"UploadRetries" json:"UploadRetries,omitempty"`
	Retention     *S3Retention `yaml:"Retention" json:"Retention,omitempty"`
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
	if t.Retention != nil {
		if t.Retention.KeepCount < 0 {
			return fmt.Errorf("S3Target Retention KeepCount must be >= 0")
		}
		if t.Retention.KeepDays < 0 {
			return fmt.Errorf("S3Target Retention KeepDays must be >= 0")
		}
		if t.Retention.KeepCount == 0 && t.Retention.KeepDays == 0 {
			return fmt.Errorf("S3Target Retention requires KeepCount or KeepDays (or both)")
		}
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
			if t.Retention != nil {
				if retErr := t.applyRetention(); retErr != nil {
					log.Warn().Err(retErr).Msg("S3 retention cleanup failed")
				}
			}
			return nil
		}
	}
	return err
}

func (t S3Target) applyRetention() error {
	objects, err := s3_utils.ListObjects(
		t.AccessKey,
		t.SecretKey,
		t.Region,
		t.Endpoint,
		t.BucketName,
		t.Prefix,
	)
	if err != nil {
		return fmt.Errorf("retention: list objects: %w", err)
	}

	// Sort by LastModified descending (newest first)
	sort.Slice(objects, func(i, j int) bool {
		return objects[i].LastModified.After(*objects[j].LastModified)
	})

	now := time.Now().UTC()
	toDelete := make(map[string]bool)

	for i, obj := range objects {
		if obj.Key == nil || obj.LastModified == nil {
			continue
		}

		deleteByCount := t.Retention.KeepCount > 0 && i >= t.Retention.KeepCount
		deleteByAge := t.Retention.KeepDays > 0 && now.Sub(*obj.LastModified) > time.Duration(t.Retention.KeepDays)*24*time.Hour

		if deleteByCount || deleteByAge {
			toDelete[*obj.Key] = true
		}
	}

	if len(toDelete) == 0 {
		return nil
	}

	keys := make([]string, 0, len(toDelete))
	for key := range toDelete {
		keys = append(keys, key)
	}

	log.Info().Int("count", len(keys)).Msg("S3 retention: deleting old backups")

	return s3_utils.DeleteObjects(
		t.AccessKey,
		t.SecretKey,
		t.Region,
		t.Endpoint,
		t.BucketName,
		keys,
	)
}
