package backup_output

import (
	"io"
)

type BackupOutput struct {
	Data   io.ReadSeeker
	Stderr string
}
