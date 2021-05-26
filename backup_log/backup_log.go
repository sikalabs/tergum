package backup_log

import (
	"bytes"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

type BackupLog struct {
	BackupID      string
	DestinationID string
	Success       bool
	Error         error
}

type BackupGlobalLog struct {
	Logs []BackupLog
}

func (gl *BackupGlobalLog) Success() bool {
	for _, log := range gl.Logs {
		if log.Success == false {
			return false
		}
	}
	return true
}

func (gl *BackupGlobalLog) SuccessString() string {
	if gl.Success() {
		return "OK"
	}
	return "ERR"
}

func renderGlobalLogTable(globalLog BackupGlobalLog, writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Success", "Backup", "Target", "Error"})

	for _, log := range globalLog.Logs {
		var strStatus, strError string

		if log.Success {
			strStatus = "OK"
		} else {
			strStatus = "ERROR"
		}

		if log.Error == nil {
			strError = ""
		} else {
			strError = log.Error.Error()
		}

		table.Append([]string{
			strStatus,
			log.BackupID,
			log.DestinationID,
			strError,
		})
	}
	table.Render()
}

func GlobalLogToString(globalLog BackupGlobalLog) string {
	buf := new(bytes.Buffer)
	renderGlobalLogTable(globalLog, buf)
	return buf.String()
}

func GlobalLogToOutput(globalLog BackupGlobalLog) {
	renderGlobalLogTable(globalLog, os.Stdout)
}
