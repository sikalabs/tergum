package symmetric_encryption

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/sikalabs/tergum/utils/rand_utils"
)

type SymmetricEncryptionMiddleware struct {
	Passphrase string `yaml:"Passphrase" json:"Passphrase,omitempty"`
}

func (m SymmetricEncryptionMiddleware) Validate() error {
	if m.Passphrase == "" {
		return fmt.Errorf("SymmetricEncryptionMiddleware need to have a Passphrase")
	}
	return nil
}

func (m SymmetricEncryptionMiddleware) Process(data io.ReadSeeker) (io.ReadSeeker, error) {
	// Encrypt:
	//     gpg --batch --output greetings.txt.gpg --passphrase mypassword --symmetric greetings.txt
	// Decrypt:
	//     gpg --batch --output greetings1.txt --passphrase mypassword --decrypt greetings.txt.gpg

	var err error

	inputFile, err := os.CreateTemp("", "tergum-symmetric-encryption-input-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(inputFile.Name())
	io.Copy(inputFile, data)

	outputFileName := path.Join(os.TempDir(), rand_utils.GetRandString(16))

	cmd := exec.Command(
		"gpg",
		"--batch",
		"--output", outputFileName,
		"--passphrase", m.Passphrase,
		"--symmetric", inputFile.Name(),
	)
	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	outputFile, err := os.Open(outputFileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	outputFile.Seek(0, 0)

	return outputFile, nil
}
