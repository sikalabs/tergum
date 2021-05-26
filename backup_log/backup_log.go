package backup_log

import (
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

func ShowGlobalLog(globalLog BackupGlobalLog) {
	table := tablewriter.NewWriter(os.Stdout)
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
