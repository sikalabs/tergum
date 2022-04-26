package azure_blob_utils

import (
	"context"
	"fmt"
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func Upload(accountName, accountKey, containerName, fileName string, data io.ReadSeeker) error {
	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return err
	}

	serviceClient, err := azblob.NewServiceClientWithSharedKey(fmt.Sprintf("https://%s.blob.core.windows.net/", accountName), cred, nil)
	if err != nil {
		return err
	}

	containerClient := serviceClient.NewContainerClient(containerName)

	blockBlobClient := containerClient.NewBlockBlobClient(fileName)

	_, err = blockBlobClient.Upload(context.TODO(), streaming.NopCloser(data), nil)
	if err != nil {
		return err
	}

	return nil
}
