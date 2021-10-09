package do_backup

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/sikalabs/tergum/backup"
)

func remoteBackup(backup backup.Backup) (io.ReadSeeker, error) {
	var err error
	json_data, err := json.Marshal(backup.Source)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(
		backup.RemoteExec.Server,
		"application/json",
		bytes.NewBuffer(json_data),
	)
	if err != nil {
		return nil, err
	}
	outputFile, err := os.CreateTemp("", "tergum-http-")
	if err != nil {
		return nil, err
	}
	io.Copy(outputFile, resp.Body)
	return outputFile, nil
}
