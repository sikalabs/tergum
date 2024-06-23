package azure_blob

import (
	"fmt"
	"io"

	"github.com/sikalabs/tergum/utils/azure_blob_utils"
	"github.com/sikalabs/tergum/utils/file_utils"
)

type AzureBlobTarget struct {
	AccountName   string `yaml:"AccountName" json:"AccountName,omitempty"`
	AccountKey    string `yaml:"AccountKey" json:"AccountKey,omitempty"`
	ContainerName string `yaml:"ContainerName" json:"ContainerName,omitempty"`
	Prefix        string `yaml:"Prefix" json:"Prefix,omitempty"`
	Suffix        string `yaml:"Suffix" json:"Suffix,omitempty"`
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
