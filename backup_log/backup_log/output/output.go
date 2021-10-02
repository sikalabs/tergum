package output

import (
	"bytes"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/sikalabs/tergum/backup_log"
)

func BackupLogTable(l backup_log.BackupLog, writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{"Success", "Backup", "Target", "Error"})

	for _, log := range l.Events {
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
			log.SourceName + ": " + log.BackupID,
			log.TargetName + ": " + log.TargetID,
			strError,
		})
	}
	table.Render()
}

func BackupLogToString(l backup_log.BackupLog) string {
	buf := new(bytes.Buffer)
	BackupLogTable(l, buf)
	return buf.String()
}

func BackupLogToOutput(l backup_log.BackupLog) {
	BackupLogTable(l, os.Stdout)
}
