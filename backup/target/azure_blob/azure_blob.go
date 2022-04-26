package azure_blob

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/utils/azure_blob_utils"
	"github.com/sikalabs/tergum/utils/file_utils"
)

type AzureBlobTarget struct {
	AccountName   string `yaml:"AccountName"`
	AccountKey    string `yaml:"AccountKey"`
	ContainerName string `yaml:"ContainerName"`
	Prefix        string `yaml:"Prefix"`
	Suffix        string `yaml:"Suffix"`
}

func (t AzureBlobTarget) Validate() error {
	if t.AccountName == "" {
		return fmt.Errorf("AzureBlobTarget need to have a AccountName")
	}
	if t.AccountKey == "" {
		return fmt.Errorf("AzureBlobTarget requires AccountKey")
	}
	if t.ContainerName == "" {
		return fmt.Errorf("AzureBlobTarget requires region or ContainerName")
	}
	if t.Prefix == "" {
		return fmt.Errorf("AzureBlobTarget requires Prefix")
	}
	if t.Suffix == "" {
		return fmt.Errorf("AzureBlobTarget requires Suffix")
	}
	return nil
}

func (t AzureBlobTarget) Save(data io.ReadSeeker) error {
	return azure_blob_utils.Upload(
		t.AccountName,
		t.AccountKey,
		t.ContainerName,
		file_utils.GetFileName(t.Prefix, t.Suffix),
		data,
	)
}
