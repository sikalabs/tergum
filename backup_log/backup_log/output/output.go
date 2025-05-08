package output

import (
	"bytes"
	"io"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/sikalabs/tergum/backup_log"
	"github.com/sikalabs/tergum/utils/file_size_utils"
)

func BackupLogTable(l backup_log.BackupLog, writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{
		"Success",
		"Backup", "Backup Time",
		"Target", "Upload Time", "File Size",
		"Error",
		"Time Total",
	})

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
			strconv.Itoa(log.BackupDuration) + "s" +
				" (+" + strconv.Itoa(log.BackupMiddlewaresDuration) + "s)",
			log.TargetName + ": " + log.TargetID,
			strconv.Itoa(log.TargetDuration) + "s" +
				" (+" + strconv.Itoa(log.TargetMiddlewaresDuration) + "s)",
			file_size_utils.PrettyFileSize(log.TargetFileSize),
			strError,
			strconv.Itoa(log.TotalDuration()) + "s",
		})
	}
	table.Render()
}

func BackupLogTelegram(l backup_log.BackupLog) string {
	out := new(bytes.Buffer)
	out.WriteString("=== log ===\n")
	out.WriteString("\n")

	for _, log := range l.Events {
		var strStatus, emojiStatus string

		if log.Success {
			strStatus = "OK"
			emojiStatus = "\u2705"
		} else {
			strStatus = "ERROR"
			emojiStatus = "\u274C"
		}

		out.WriteString("Success: " + strStatus + " " + emojiStatus + "\n")
		out.WriteString("Backup: " + log.SourceName + ": " + log.BackupID + "\n")
		out.WriteString("Backup Time: " + strconv.Itoa(log.BackupDuration) + "s" +
			" (+" + strconv.Itoa(log.BackupMiddlewaresDuration) + "s)" + "\n")
		out.WriteString("Target: " + log.TargetName + ": " + log.TargetID + "\n")
		out.WriteString("Upload Time: " + strconv.Itoa(log.TargetDuration) + "s" +
			" (+" + strconv.Itoa(log.TargetMiddlewaresDuration) + "s)" + "\n")
		out.WriteString("File Size: " + file_size_utils.PrettyFileSize(log.TargetFileSize) + "\n")
		if log.Error != nil {
			out.WriteString("Error: " + log.Error.Error() + "\n")
		}
		out.WriteString("Time Total: " + strconv.Itoa(log.TotalDuration()) + "s" + "\n")
		out.WriteString("\n")
	}
	return out.String()
}

func BackupLogToString(l backup_log.BackupLog) string {
	buf := new(bytes.Buffer)
	BackupLogTable(l, buf)
	return buf.String()
}

func BackupLogToTelegramString(l backup_log.BackupLog) string {
	return BackupLogTelegram(l)
}

func BackupLogToOutput(l backup_log.BackupLog) {
	BackupLogTable(l, os.Stdout)
}

func BackupErrorLogTable(l backup_log.BackupLog, writer io.Writer) {
	table := tablewriter.NewWriter(writer)
	table.SetHeader([]string{
		"Backup",
		"Target",
		"Error",
	})

	for _, log := range l.Events {
		if log.Success {
			continue
		}

		table.Append([]string{
			log.BackupID,
			log.TargetID,
			log.StdErr,
		})
	}
	table.Render()
}

func BackupErrorLogTelegram(l backup_log.BackupLog) string {
	noErrors := true
	out := new(bytes.Buffer)
	out.WriteString("=== errors ===\n")
	out.WriteString("\n")
	for _, log := range l.Events {
		if log.Success {
			continue
		}
		noErrors = false
		out.WriteString("Backup: " + log.BackupID + "\n")
		out.WriteString("Target: " + log.TargetID + "\n")
		out.WriteString("Error: " + log.StdErr + "\n")
		out.WriteString("\n")
	}
	if noErrors {
		return ""
	}
	return out.String()
}

func BackupErrorLogToString(l backup_log.BackupLog) string {
	buf := new(bytes.Buffer)
	BackupErrorLogTable(l, buf)
	return buf.String()
}

func BackupErrorLogToTelegramString(l backup_log.BackupLog) string {
	return BackupErrorLogTelegram(l)
}

func BackupErrorLogToOutput(l backup_log.BackupLog) {
	BackupErrorLogTable(l, os.Stdout)
}
