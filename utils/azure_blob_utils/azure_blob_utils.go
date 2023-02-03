package azure_blob_utils

import (
	"context"
	"fmt"
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

func Upload(accountName, accountKey, containerName, fileName string, data io.ReadSeeker) error {
	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return err
	}

	containerClient, err := container.NewClientWithSharedKeyCredential(fmt.Sprintf("https://%s.blob.core.windows.net/", accountName), cred, nil)
	if err != nil {
		return err
	}

	blockBlobClient := containerClient.NewBlockBlobClient(fileName)

	_, err = blockBlobClient.Upload(context.TODO(), streaming.NopCloser(data), nil)
	if err != nil {
		return err
	}

	return nil
}
